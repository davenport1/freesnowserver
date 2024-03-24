package api

import (
	"encoding/json"
	"net/http"
)

// HealthCheck returns the status of the server
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status": "available",
	}

	jsonHealth, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jsonHealth = append(jsonHealth, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonHealth)
}
