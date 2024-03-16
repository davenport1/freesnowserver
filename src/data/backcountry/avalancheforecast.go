package backcountry

import (
	"database/sql"
	"time"
)

// AvalancheForecast represents all the different factors of the avalanche forecast
// ID links to weather forecast for the zone forecasted and the different models that make the forecast
type AvalancheForecast struct {
	ID           int       `json:"id"`
	ForecastDate time.Time `json:"forecastDate"`
	CreatedAt    time.Time `json:"createdAt"`
}

type AvalancheForecastModel struct {
	Db *sql.DB
}
