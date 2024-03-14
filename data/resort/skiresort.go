package resort

import (
	"database/sql"
	"freesnow/common"
	"freesnow/data/weather"
	"time"
)

type SkiResort struct {
	ID              int64              `json:"id"`
	ResortName      string             `json:"resortName"`
	Coordinates     common.Coordinates `json:"coordinates"`
	CreatedAt       time.Time          `json:"createdAt"`
	SnowReport      SnowReport         `json:"snowReport"`
	WeatherForecast weather.Forecast   `json:"weatherForecast"`
	TrailReport     TrailReport        `json:"trailReport"`
	LiftReport      LiftReport         `json:"liftReport"`
	Version         int32              `json:"version"`
}

type SkiResortModel struct {
	Db *sql.DB
}

// InsertNewResort - Creates a new resort with the given resort.
func (s SkiResortModel) InsertNewResort(resort *SkiResort) error {
	resort.CreatedAt = time.Now().UTC()
	timezone := "PST"
	query := `INSERT INTO ski_resort (resort_name, created_at, timezone, version)
				VALUES ($1, $2, $3, $4)`

	args := []interface{}{
		resort.ResortName,
	}

	return s.Db.QueryRow(query, args...).Scan(&resort.ID, &resort.CreatedAt, timezone, &resort.Version)
}

//func (s SkiResortModel) UpdateResort(resort *SkiResort) error {
//
//}
