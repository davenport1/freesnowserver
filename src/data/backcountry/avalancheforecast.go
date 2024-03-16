package backcountry

import "time"

type AvalancheForecast struct {
	ID           int       `json:"id"`
	ForecastDate time.Time `json:"forecastDate"`
}
