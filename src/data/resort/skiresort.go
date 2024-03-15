package resort

import (
	"database/sql"
	"errors"
	"fmt"
	"freesnow/data/coordinates"
	"freesnow/data/weather"
	"time"
)

type SkiResort struct {
	ID              int64                   `json:"id"`
	ResortName      string                  `json:"resortName"`
	Coordinates     coordinates.Coordinates `json:"coordinates"`
	CreatedAt       time.Time               `json:"createdAt"`
	LastUpdated     time.Time               `json:"lastUpdated"`
	SnowReport      SnowReport              `json:"snowReport"`
	WeatherForecast weather.Forecast        `json:"weatherForecast"`
	TrailReport     TrailReport             `json:"trailReport"`
	LiftReport      LiftReport              `json:"liftReport"`
	Version         int32                   `json:"version"`
}

type SkiResortModel struct {
	Db *sql.DB
}

// InsertNewResort - Creates a new resort with the given resort.
func (s SkiResortModel) InsertNewResort(resort *SkiResort) error {
	resort.CreatedAt = time.Now().UTC()
	timezone := "PST"
	query := `
		INSERT INTO ski_resort (resort_name, created_at, timezone)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version`

	args := []interface{}{
		resort.ResortName,
		time.Now().UTC(),
		timezone,
	}

	return s.Db.QueryRow(query, args...).Scan(&resort.ID, &resort.CreatedAt, &resort.Version)
}

func (s SkiResortModel) DeleteResort(id int64) error {
	query := `
 	DELETE FROM ski_resort
	WHERE id = $1`

	results, err := s.Db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (s SkiResortModel) GetSkiResortByName(name string) (*SkiResort, error) {
	// handle empty name at controller, validation at middleware
	query := `
	SELECT id, resort_name, created_at, version 
	FROM ski_resort 
	WHERE resort_name LIKE '%$1%'
	ORDER BY created_at DESC 
	LIMIT 1`

	var skiResort SkiResort
	if err := s.Db.QueryRow(query, name).Scan(
		&skiResort.ID,
		&skiResort.ResortName,
		&skiResort.CreatedAt,
		&skiResort.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &skiResort, nil
}

func (s SkiResortModel) GetAllResorts() ([]*SkiResort, error) {
	query := `
	SELECT id, resort_name, created_at, version
	FROM ski_resort ORDER BY resort_name`
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Failed to close rows")
		}
	}(rows)

	resorts := []*SkiResort{}

	for rows.Next() {
		var skiResort SkiResort
		if err := rows.Scan(
			&skiResort.ID,
			&skiResort.ResortName,
			&skiResort.CreatedAt,
			&skiResort.Version,
		); err != nil {
			return nil, err
		}

		resorts = append(resorts, &skiResort)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resorts, nil
}
