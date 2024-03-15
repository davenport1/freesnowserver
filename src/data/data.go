package data

import (
	"database/sql"
	"freesnow/data/resort"
)

// Models - struct containing database entities and the database connection
type Models struct {
	DB         *sql.DB
	SkiResorts resort.SkiResortModel
}

// NewModels - Creates a Models struct with the database connection and database entities
func NewModels(db *sql.DB) Models {
	return Models{
		DB:         db,
		SkiResorts: resort.SkiResortModel{Db: db},
	}
}
