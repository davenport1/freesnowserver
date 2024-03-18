package weather

import (
	"database/sql"
	"freesnow/common"
	"time"
)

type WeatherForecast struct {
	ID                 int64                `json:"id"`
	TempHigh           int64                `json:"temperatureHigh"`
	TempLow            int64                `json:"temperatureLow"`
	WindDirection      int64                `json:"windDirection"`
	WindSpeed          int64                `json:"windSpeed"`
	OvercastLevel      common.OvercastLevel `json:"overcastLevel"`
	HumidityPercentage int64                `json:"humidityPercentage"`
	TempFeelsLike      int64                `json:"temperatureFeelsLike"`
	TempWindChill      int64                `json:"temperatureWithWindChill"`
	Sunrise            time.Time            `json:"sunriseTime"`
	Sunset             time.Time            `json:"sunsetTime"`
}

type ForecastResort struct {
	ID              int64           `json:"id"`
	SkiResortId     int64           `json:"skiResortId"`
	WeatherForecast WeatherForecast `json:"weatherForecast"`
}

type ForecastBackcountry struct {
	ID                int64           `json:"id"`
	BackcountryZoneId int64           `json:"backcountryZoneId"`
	WeatherForecast   WeatherForecast `json:"weatherForecast"`
}

type ForecastGeneral struct {
	ID                 int64                `json:"id"`
	TempHigh           int64                `json:"temperatureHigh"`
	TempLow            int64                `json:"temperatureLow"`
	WindDirection      int64                `json:"windDirection"`
	WindSpeed          int64                `json:"windSpeed"`
	OvercastLevel      common.OvercastLevel `json:"overcastLevel"`
	HumidityPercentage int64                `json:"humidityPercentage"`
	TempFeelsLike      int64                `json:"temperatureFeelsLike"`
	TempWindChill      int64                `json:"temperatureWithWindChill"`
	Sunrise            time.Time            `json:"sunriseTime"`
	Sunset             time.Time            `json:"sunsetTime"`
	Location           time.Location        `json:"location"`
}

type WeatherModel struct {
	Db *sql.DB
}
