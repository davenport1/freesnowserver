package api

import (
	"freesnow/common"
	"freesnow/data"
	"freesnow/data/backcountry"
	"net/http"
	"strconv"
	"time"
)

type forecastRequestModel struct {
	ZoneId              int                   `json:"zoneId"`
	ForecastDate        time.Time             `json:"forecastDate"`
	BottomLine          string                `json:"bottomLine"`
	OverallDanger       int                   `json:"overallDanger"`
	DangerAboveTreeline int                   `json:"dangerAboveTreeline"`
	DangerAtTreeLine    int                   `json:"dangerAtTreeLine"`
	DangerBelowTreeline int                   `json:"dangerBelowTreeline"`
	AvalancheProblems   []problemRequestModel `json:"avalancheProblems"`
}

type problemRequestModel struct {
	AdditionalNotes string `json:"additionalNotes"`
	Aspects         int    `json:"aspect"`
	Elevations      int    `json:"elevation"`
	ProblemType     int    `json:"problemType"`
	Likelihood      int    `json:"likelihood"`
	Size            int    `json:"size"`
}

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
func SaveNewForecast(w http.ResponseWriter, r *http.Request, m *data.Models) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	var forecastRequest forecastRequestModel

	if err := readJSON(w, r, &forecastRequest); err != nil {
		// log
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	var avalancheProblems []backcountry.AvalancheProblem

	for _, problemRequest := range forecastRequest.AvalancheProblems {
		newProblem := backcountry.AvalancheProblem{
			AdditionalNotes: problemRequest.AdditionalNotes,
			Aspect:          common.AvalancheAspect(problemRequest.Aspects),
			Elevation:       common.AvalancheElevation(problemRequest.Elevations),
			ProblemType:     common.AvalancheProblemType(problemRequest.ProblemType),
			Likelihood:      common.AvalancheLikelihood(problemRequest.Likelihood),
			Size:            common.AvalancheSize(problemRequest.Size),
		}

		avalancheProblems = append(avalancheProblems, newProblem)
	}

	newForecast := &backcountry.AvalancheForecast{
		ForecastZoneId:      forecastRequest.ZoneId,
		ForecastDate:        forecastRequest.ForecastDate,
		BottomLine:          forecastRequest.BottomLine,
		OverallDanger:       common.AvalancheDanger(forecastRequest.OverallDanger),
		DangerAboveTreeline: common.AvalancheDanger(forecastRequest.DangerAboveTreeline),
		DangerAtTreeline:    common.AvalancheDanger(forecastRequest.DangerAtTreeLine),
		DangerBelowTreeline: common.AvalancheDanger(forecastRequest.DangerBelowTreeline),
		AvalancheProblems:   avalancheProblems,
	}

	if err := m.AvalancheForecasts.SaveNewForecast(newForecast); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	// set headers?
	if err := writeJSON(w,
		int(http.StatusCreated),
		envelope{"AvalancheForecast": newForecast},
		nil,
	); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// GetArchivedForecasts - gets forecasts that have been archived in the server.
func GetArchivedForecasts(w http.ResponseWriter, r *http.Request) {

}

// RemoveArchivedForecasts - removes all archived forecasts from the server
func RemoveArchivedForecasts(w http.ResponseWriter, r *http.Request) {

}

// GetRecentForecastByZone - Gets the forecast for the backcountry zone specified by the external id requested.
//
// externalId - references the id used by the avalanche centers for their api
//
// GET api/v1/avalanche/zone/forecast/{externalId}
func GetRecentForecastByZone(w http.ResponseWriter, r *http.Request, m *data.Models) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	externalIdStr := r.URL.Path[len("/v1/avalanche/zone/forecast/"):]
	externalId, err := strconv.ParseInt(externalIdStr, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}
	if !m.ForecastZones.ZoneExistsByExternalId(int(externalId)) {

	}
}
