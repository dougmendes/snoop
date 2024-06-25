package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dougmendes/snoopy/model"
)

func ReadJSON(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("test.json")
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
