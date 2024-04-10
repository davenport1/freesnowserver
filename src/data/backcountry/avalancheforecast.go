package backcountry

import (
	"database/sql"
	"errors"
	"fmt"
	"freesnow/common"
	"freesnow/data/weather"
	"time"
)

// AvalancheForecast represents all the different factors of the avalanche forecast
// ID links to weather forecast for the zone forecasted and the different models that make the forecast
type AvalancheForecast struct {
	ID                  int                         `json:"id"`
	ForecastZoneId      int                         `json:"forecastZoneId"`
	ForecastDate        time.Time                   `json:"forecastDate"`
	CreatedAt           time.Time                   `json:"createdAt"`
	BottomLine          string                      `json:"bottomLine"`
	OverallDanger       common.AvalancheDanger      `json:"overallDanger"`
	DangerAboveTreeline common.AvalancheDanger      `json:"dangerAboveTreeline"`
	DangerAtTreeline    common.AvalancheDanger      `json:"dangerAtTreeline"`
	DangerBelowTreeline common.AvalancheDanger      `json:"dangerBelowTreeline"`
	AvalancheProblems   []*AvalancheProblem         `json:"avalancheProblems"`
	CurrentWeather      weather.ForecastBackcountry `json:"currentWeather"`
}

// AvalancheProblem represents individual avalanche problems. Avalanche forecasts can (and usually do)
// have multiple of these
type AvalancheProblem struct {
	ID                  int                         `json:"id"`
	AdditionalNotes     string                      `json:"additionalNotes"`
	AvalancheForecastId int                         `json:"avalancheForecastId"`
	Aspect              common.AvalancheAspect      `json:"aspect"`
	Elevation           common.AvalancheElevation   `json:"elevation"`
	ProblemType         common.AvalancheProblemType `json:"problemType"`
	Likelihood          common.AvalancheLikelihood  `json:"likelihood"`
	Size                common.AvalancheSize        `json:"size"`
}

// AvalancheForecastModel provides the connection between the models and database
type AvalancheForecastModel struct {
	Db *sql.DB
}

func (a AvalancheForecastModel) SaveNewForecast(forecast *AvalancheForecast) error {
	forecastQuery := `
	INSERT INTO avalanche_forecast (
		forecast_zone_id, 
		forecast_date,
		bottom_line, 
		overall_danger, 
		danger_above_tl,
		danger_at_tl, 
		danger_below_tl)
	VALUES ($1, $2, $3, $4, $5, $6, &7)
	RETURNING id`

	problemQuery := `
	INSERT INTO avalanche_problem (
	   avalanche_forecast_id,
	   additional_notes, 
	   aspect, 
	   elevation, 
	   problem_type,
	   likelihood, 
	   likely_size)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	forecastArgs := []interface{}{
		forecast.ForecastZoneId,
		forecast.ForecastDate,
		forecast.BottomLine,
		forecast.OverallDanger,
		forecast.DangerAboveTreeline,
		forecast.DangerAtTreeline,
		forecast.DangerBelowTreeline,
	}

	if err := a.Db.QueryRow(forecastQuery, forecastArgs...).Scan(&forecast.ID); err != nil {
		return err
	}

	for _, problem := range forecast.AvalancheProblems {
		problem.AvalancheForecastId = forecast.ID
		problemArgs := []interface{}{
			problem.AvalancheForecastId,
			problem.AdditionalNotes,
			problem.Aspect,
			problem.Elevation,
			problem.ProblemType,
			problem.Likelihood,
			problem.Size,
		}
		if err := a.Db.QueryRow(problemQuery, problemArgs...).Scan(&problem.ID); err != nil {
			return err
		}
	}

	return nil
}

func (a AvalancheForecastModel) GetCurrentForecastForZone(zoneId int) (*AvalancheForecast, error) {
	query := `
	SELECT * 
	FROM avalanche_forecast 
	WHERE forecast_zone_id = $1 
	ORDER BY forecast_date DESC 
	LIMIT 1`

	var forecast AvalancheForecast
	if err := a.Db.QueryRow(query, zoneId).Scan(
		&forecast.ID,
		&forecast.ForecastZoneId,
		&forecast.ForecastDate,
		&forecast.CreatedAt,
		&forecast.BottomLine,
		&forecast.OverallDanger,
		&forecast.DangerAboveTreeline,
		&forecast.DangerAtTreeline,
		&forecast.DangerBelowTreeline); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("no forecasts exist for the zone id")
		default:
			return nil, err
		}
	}

	problems, err := a.GetAvalancheProblemsForForecastId(forecast.ID)
	if err != nil {
		return nil, err
	}

	for _, problem := range problems {
		forecast.AvalancheProblems = append(forecast.AvalancheProblems, problem)
	}

	return &forecast, nil
}

func (a AvalancheForecastModel) GetAvalancheProblemsForForecastId(forecastId int) ([]*AvalancheProblem, error) {
	query := `
	SELECT * FROM avalanche_problem 
	WHERE avalanche_forecast_id = $1`

	var problems []*AvalancheProblem
	rows, err := a.Db.Query(query, forecastId)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			fmt.Println("failed to close rows")
		}
	}(rows)

	for rows.Next() {
		var problem AvalancheProblem
		if err := rows.Scan(
			&problem.ID,
			&problem.AvalancheForecastId,
			&problem.AdditionalNotes,
			&problem.Aspect,
			&problem.Elevation,
			&problem.ProblemType,
			&problem.Likelihood,
			&problem.Size,
		); err != nil {
			return nil, err
		}

		problems = append(problems, &problem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return problems, nil
}
