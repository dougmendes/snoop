package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dougmendes/snoopy/model"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON_FileOpenError(t *testing.T) {
	mockOpen := func(name string) (*os.File, error) {
		return nil, errors.New("mock error")
	}

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReadJSONWithFileOpener(w, r, mockOpen)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "mock error\n", rr.Body.String())
}

func TestReadJSON_DecodeError(t *testing.T) {
	mockOpen := func(name string) (*os.File, error) {
		file, err := os.CreateTemp("", "invalid*.json")
		if err != nil {
			return nil, err
		}
		file.WriteString("invalid json")
		file.Seek(0, 0) // Reset file pointer to beginning
		return file, nil
	}

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReadJSONWithFileOpener(w, r, mockOpen)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid character")
}

func TestReadJSON_Success(t *testing.T) {
	expectedData := []model.ScanResult{
		{
			Target: "target1",
			Vulnerability: []model.Vulnerability{
				{
					VulnerabilityID:  "vuln1",
					PkgName:          "pkg1",
					InstalledVersion: "1.0",
					FixedVersion:     "1.1",
					Severity:         "high",
					Description:      "desc1",
					References:       []string{"ref1", "ref2"},
				},
			},
		},
	}

	tempFile, err := os.CreateTemp("", "test*.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	jsonData, err := json.Marshal(expectedData)
	assert.NoError(t, err)

	tempFile.Write(jsonData)
	tempFile.Seek(0, 0) // Reset file pointer to beginning

	mockOpen := func(name string) (*os.File, error) {
		return tempFile, nil
	}

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReadJSONWithFileOpener(w, r, mockOpen)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var actualData []model.ScanResult
	err = json.NewDecoder(rr.Body).Decode(&actualData)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)
}
