package api

import (
	"freesnow/data"
	"net/http"
)

// GetAllResorts retrieves all resorts and its associated data
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

// Resort provides routing for endpoint /v1/resort filtering by the http method
func Resort(w http.ResponseWriter, r *http.Request, m *data.Models) {
	switch r.Method {
	case http.MethodGet:
		GetResort(w, r, m)
	case http.MethodPost:
		SaveNewResort(w, r, m)
	case http.MethodPut:
		UpdateResort(w, r, m)
	case http.MethodDelete:
		DeleteResort(w, r, m)
	}

}

// SaveNewResort Saves a new resort and returns a status code to the client
func SaveNewResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

// GetResort Retrieves the resort queried by the client, or 404 if none found
func GetResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

// UpdateResort updates the resort with new information
func UpdateResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}

// DeleteResort removes a resort from the database and returns a status code to the client
func DeleteResort(w http.ResponseWriter, r *http.Request, m *data.Models) {

}
