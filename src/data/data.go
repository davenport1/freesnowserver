package data

import (
	"database/sql"
	"freesnow/data/backcountry"
	"freesnow/data/resort"
)

// Models - struct containing database entities and the database connection
type Models struct {
	DB                 *sql.DB
	SkiResorts         resort.SkiResortModel
	AvalancheForecasts backcountry.AvalancheForecastModel
	ForecastZones      backcountry.ForecastZoneModel
}

// NewModels - Creates a Models struct with the database connection and database entities
func NewModels(db *sql.DB) Models {
	return Models{
		DB:                 db,
		SkiResorts:         resort.SkiResortModel{Db: db},
		AvalancheForecasts: backcountry.AvalancheForecastModel{Db: db},
		ForecastZones:      backcountry.ForecastZoneModel{Db: db},
	}
}
