package data

import (
	"database/sql"
	"freesnow/common"
	"time"
)

type SkiResort struct {
	ID              int64
	ResortName      string
	Coordinates     common.Coordinates
	CreatedAt       time.Time
	SnowReport      SnowReport
	WeatherForecast WeatherForecast
	TrailReport     TrailReport
	LiftReport      LiftReport
	Version         int32
}

type SnowReport struct {
}

type WeatherForecast struct {
}

type TrailReport struct {
}

type LiftReport struct {
}

type SkiResortModel struct {
	Db *sql.DB
}

// InsertNewResort - Creates a new resort with the given resort.
func (s SkiResortModel) InsertNewResort(resort *SkiResort) error {
	query := `SOME SQL`

	args := []interface{}{
		resort.ResortName,
	}

	return s.Db.QueryRow(query, args...).Scan(&resort.ID, &resort.CreatedAt, &resort.Version)
}
