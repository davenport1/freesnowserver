package backcountry

import (
	"database/sql"
	"freesnow/data/weather"
	"time"
)

// ForecastZone - represents the zone that gets avalanche forecasts in terms of location, dates it has been updated,
// and provides an ID to link avalanche and weather forecasts for the zone
type ForecastZone struct {
	ID              int64                       `json:"id"`
	CreatedAt       time.Time                   `json:"createdAt"`
	LastUpdated     time.Time                   `json:"lastUpdated"`
	CurrentForecast AvalancheForecast           `json:"currentForecast"`
	CurrentWeather  weather.ForecastBackcountry `json:"currentWeather"`
}

// ForecastZoneModel - uses the db connection for CRUD
type ForecastZoneModel struct {
	Db *sql.DB
}
