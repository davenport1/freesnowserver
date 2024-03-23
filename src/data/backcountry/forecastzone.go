package backcountry

import (
	"database/sql"
	"freesnow/common"
	"freesnow/data/weather"
	"time"
)

// ForecastStation is the station that serves the zones within its boundaries avalanche forecasts.
// ForecastStations may serve multiple avalanche forecasts across different avalanche zones
type ForecastStation struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdated   time.Time `json:"lastUpdated"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Timezone      string    `json:"timezone"`
	NumberOfZones int64     `json:"numberOfZones"`
}

// ForecastZone - represents the zone that gets avalanche forecasts in terms of location, dates it has been updated,
// and provides an ID to link avalanche and weather forecasts for the zone
type ForecastZone struct {
	ID                int64                       `json:"id"`
	ForecastStationId int64                       `json:"forecastStationId"`
	Name              string                      `json:"name"`
	CreatedAt         time.Time                   `json:"createdAt"`
	LastUpdated       time.Time                   `json:"lastUpdated"`
	Latitude          float64                     `json:"latitude"`
	Longitude         float64                     `json:"longitude"`
	Timezone          string                      `json:"timezone"`
	CurrentForecast   AvalancheForecast           `json:"currentForecast"`
	CurrentWeather    weather.ForecastBackcountry `json:"currentWeather"`
}

// ForecastZoneModel - uses the db connection for CRUD
type ForecastZoneModel struct {
	Db *sql.DB
}

func (f ForecastZoneModel) SaveNewForecastStation(station *ForecastStation) error {
	station.CreatedAt = time.Now().UTC()

	query := `
 	INSERT INTO forecast_station (station_name, created_at, last_updated, time_zone)
 	VALUES ($1. $2, $3, $4)
 	RETURNING id`

	tz, err := common.GetTimeZone(station.Latitude, station.Longitude)
	if err != nil {
		return err
	}

	station.Timezone = tz
	args := []interface{}{
		station.Name,
		station.CreatedAt,
		station.LastUpdated,
		station.Timezone,
	}

	return f.Db.QueryRow(query, args...).Scan(&station.ID)
}

// SaveNewForecastZone saves the new forecast zone from the request, applying the timezone
// given the coordinates in the request. Returns the id
func (f ForecastZoneModel) SaveNewForecastZone(zone *ForecastZone) error {
	zone.CreatedAt = time.Now().UTC()

	query := `
	INSERT INTO forecast_zone (zone_name, forecast_station_id, created_at, last_updated, time_zone)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id`

	tz, err := common.GetTimeZone(zone.Latitude, zone.Longitude)
	if err != nil {
		return err
	}

	args := []interface{}{
		zone.Name,
		zone.CreatedAt,
		zone.CreatedAt,
		zone.ForecastStationId,
		tz,
	}

	return f.Db.QueryRow(query, args...).Scan(&zone.ID)
}

// GetForecastZoneByName gets the forecast zone that most closely matches the name passed,
// or an error if none are found.
// Takes the zoneName argument as a string
func (f ForecastZoneModel) GetForecastZoneByName(zoneName string) (*ForecastZone, error) {
	return nil, nil
}

// GetForecastZoneById gets the forecast zone by the id passed, or an error if none are found.
// Takes the id as an integer
func (f ForecastZoneModel) GetForecastZoneById(id int) (*ForecastZone, error) {
	return nil, nil
}

// GetAllForecastZones returns all the forecast zones currently stored in the database
func (f ForecastZoneModel) GetAllForecastZones() (*[]ForecastZone, error) {
	return nil, nil
}
