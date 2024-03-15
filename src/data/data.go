package data

import "database/sql"

// Models - struct containing database entities and the database connection
type Models struct {
	DB    *sql.DB
	Books BookModel
}

// NewModels - Creates a Models struct with the database connection and database entities
func NewModels(db *sql.DB) Models {
	return Models{
		Books: BookModel{DB: db},
		DB:    db,
	}
}
