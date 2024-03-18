package data

import (
	"database/sql"
	"freesnow/data/backcountry"
	"freesnow/data/resort"
	"freesnow/data/weather"
)

// Models - struct containing database entities and the database connection
type Models struct {
	DB                 *sql.DB
	SkiResorts         resort.SkiResortModel
	SnowReports        resort.SnowReportModel
	LiftReports        resort.LiftReportModel
	TrailReports       resort.TrailReportModel
	AvalancheForecasts backcountry.AvalancheForecastModel
	ForecastZones      backcountry.ForecastZoneModel
	WeatherForecasts   weather.WeatherModel
}

// NewModels - Creates a Models struct with the database connection and database entities
func NewModels(db *sql.DB) Models {
	return Models{
		DB:                 db,
		SkiResorts:         resort.SkiResortModel{Db: db},
		SnowReports:        resort.SnowReportModel{Db: db},
		LiftReports:        resort.LiftReportModel{Db: db},
		TrailReports:       resort.TrailReportModel{Db: db},
		AvalancheForecasts: backcountry.AvalancheForecastModel{Db: db},
		ForecastZones:      backcountry.ForecastZoneModel{Db: db},
		WeatherForecasts:   weather.WeatherModel{Db: db},
	}
}
