package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dougmendes/snoopy/model"
)

type fileOpener func(name string) (*os.File, error)

func ReadJSONWithFileOpener(w http.ResponseWriter, r *http.Request, openFile fileOpener) {
	file, err := openFile("test.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	var data []model.ScanResult
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReadJSON(w http.ResponseWriter, r *http.Request) {
	ReadJSONWithFileOpener(w, r, os.Open)
}
