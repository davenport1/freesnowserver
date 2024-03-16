package api

import "net/http"

// GetAllForecasts - More of a troubleshooting/testing endpoint. Possible will be used for ML
func GetAllForecasts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

}

// GetForecastForCoordinates - Gets the avalanche forecast for the backcountry zone that includes the coordinates given
// If the coordinates given are not in a forecasted zone, it will return forecasts for the closest areas geographically
func GetForecastForCoordinates(w http.ResponseWriter, r *http.Request) {

}

// SaveNewForecast - saves a new avalanche forecast from the serverless function
func SaveNewForecast(w http.ResponseWriter, r *http.Request) {

}

// GetArchivedForecasts - gets forecasts that have been archived in the server.
func GetArchivedForecasts(w http.ResponseWriter, r *http.Request) {

}

// RemoveArchivedForecasts - removes all archived forecasts from the server
func RemoveArchivedForecasts(w http.ResponseWriter, r *http.Request) {

}

// GetForecastByZone - Gets the forecast for the backcountry zone specified
func GetForecastByZone(w http.ResponseWriter, r *http.Request) {

}
