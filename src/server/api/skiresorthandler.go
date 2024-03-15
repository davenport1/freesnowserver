package api

import (
	"freesnow/data"
	"net/http"
)

func GetAllResorts(w http.ResponseWriter, r *http.Request, models *data.Models) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	resorts, err := models.SkiResorts.GetAllResorts()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, 200, envelope{"resorts": resorts}, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func Resort(w http.ResponseWriter, r *http.Request, m *data.Models) {
	switch r.Method {
	case http.MethodGet:
		GetResort(w, r, m)
	case http.MethodPost:
		SaveResort(w, r, m)
	case http.MethodPut:
		UpdateResort(w, r, m)
	case http.MethodDelete:
		DeleteResort(w, r, m)
	}

}

func SaveResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

func GetResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

func UpdateResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

func DeleteResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}
