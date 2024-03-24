package backcountry

import (
	"database/sql"
	"errors"
	"fmt"
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
	ExternalId    int64     `json:"externalId"`
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
	ExternalId        int64                       `json:"externalId"`
}

// ForecastZoneModel - uses the db connection for CRUD
type ForecastZoneModel struct {
	Db *sql.DB
}

func (f ForecastZoneModel) SaveNewForecastStation(station *ForecastStation) error {
	station.CreatedAt = time.Now().UTC()

	query := `
 	INSERT INTO forecast_station (station_name, created_at, last_updated, timezone)
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
	INSERT INTO forecast_zone (zone_name, forecast_station_id, created_at, last_updated, timezone)
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
	query := `
	SELECT 
	    id, 
	    forecast_station_id, 
	    zone_name, 
	    created_at,
	    last_updated,
	    ST_X(location),
	    ST_Y(location),
	    timezone,
	    external_id
	FROM forecast_zone 
	WHERE zone_name like '%$1%'
	ORDER BY created_at DESC
	LIMIT 1;`

	var zone ForecastZone
	if err := f.Db.QueryRow(query, zoneName).Scan(
		&zone.ID,
		&zone.ForecastStationId,
		&zone.Name,
		&zone.CreatedAt,
		&zone.LastUpdated,
		&zone.Longitude,
		&zone.Latitude,
		&zone.Timezone,
		&zone.ExternalId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record was not found for the given zone name")
		default:
			return nil, err
		}
	}

	return &zone, nil
}

// GetForecastZoneById gets the forecast zone by the id passed, or an error if none are found.
// Takes the id as an integer
func (f ForecastZoneModel) GetForecastZoneById(id int) (*ForecastZone, error) {
	query := `
	SELECT 
	    id, 
	    forecast_station_id, 
	    zone_name, 
	    created_at,
	    last_updated,
	    ST_X(location),
	    ST_Y(location),
	    timezone,
	    external_id
	FROM forecast_zone 
	WHERE id = $1
	ORDER BY created_at DESC
	LIMIT 1;`

	var zone ForecastZone
	if err := f.Db.QueryRow(query, id).Scan(
		&zone.ID,
		&zone.ForecastStationId,
		&zone.Name,
		&zone.CreatedAt,
		&zone.LastUpdated,
		&zone.Longitude,
		&zone.Latitude,
		&zone.Timezone,
		&zone.ExternalId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record was not found for the given zone name")
		default:
			return nil, err
		}
	}

	return &zone, nil
}

// GetAllForecastZones returns all the forecast zones currently stored in the database
func (f ForecastZoneModel) GetAllForecastZones() ([]*ForecastZone, error) {
	query := `
	SELECT * 
	FROM forecast_zone
	ORDER BY zone_name DESC`

	rows, err := f.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Failed to close rows")
		}
	}(rows)

	var zones []*ForecastZone

	for rows.Next() {
		var zone ForecastZone

		if err := rows.Scan(
			&zone.ID,
			&zone.ForecastStationId,
			&zone.Name,
			&zone.CreatedAt,
			&zone.LastUpdated,
			&zone.Longitude,
			&zone.Latitude,
			&zone.Timezone,
			&zone.ExternalId,
		); err != nil {
			return nil, err
		}

		zones = append(zones, &zone)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return zones, nil
}

func (f ForecastZoneModel) ZoneExistsByExternalId(externalId int) bool {
	query := `
 	SELECT COUNT(id)
	FROM forecast_zone 
	WHERE external_id = $1`

	args := []interface{}{
		externalId,
	}

	var count int
	if err := f.Db.QueryRow(query, args...).Scan(count); err != nil {
		return false
	}

	return count > 0
}
