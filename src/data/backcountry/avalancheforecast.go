package backcountry

import (
	"database/sql"
	"freesnow/common"
	"freesnow/data/weather"
	"time"
)

// AvalancheForecast represents all the different factors of the avalanche forecast
// ID links to weather forecast for the zone forecasted and the different models that make the forecast
type AvalancheForecast struct {
	ID                  int                         `json:"id"`
	ForecastZoneId      int                         `json:"forecastZoneId"`
	ForecastDate        time.Time                   `json:"forecastDate"`
	CreatedAt           time.Time                   `json:"createdAt"`
	BottomLine          string                      `json:"bottomLine"`
	OverallDanger       common.AvalancheDanger      `json:"overallDanger"`
	DangerAboveTreeline common.AvalancheDanger      `json:"dangerAboveTreeline"`
	DangerAtTreeline    common.AvalancheDanger      `json:"dangerAtTreeline"`
	DangerBelowTreeline common.AvalancheDanger      `json:"dangerBelowTreeline"`
	AvalancheProblems   []AvalancheProblem          `json:"avalancheProblems"`
	CurrentWeather      weather.ForecastBackcountry `json:"currentWeather"`
}

type AvalancheProblem struct {
	ID                  int                         `json:"id"`
	AdditionalNotes     string                      `json:"additionalNotes"`
	AvalancheForecastId int                         `json:"avalancheForecastId"`
	Aspect              common.AvalancheAspect      `json:"aspect"`
	Elevation           common.AvalancheElevation   `json:"elevation"`
	ProblemType         common.AvalancheProblemType `json:"problemType"`
	Likelihood          common.AvalancheLikelihood  `json:"likelihood"`
	Size                common.AvalancheSize        `json:"size"`
}

type AvalancheForecastModel struct {
	Db *sql.DB
}
