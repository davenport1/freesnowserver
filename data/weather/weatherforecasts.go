package weather

import "time"

type Forecast struct {
	ID                 int64     `json:"id"`
	SkiResortId        int64     `json:"skiResortId"`
	TempHigh           int64     `json:"temperatureHigh"`
	TempLow            int64     `json:"temperatureLow"`
	WindDirection      int64     `json:"windDirection"`
	WindSpeed          int64     `json:"windSpeed"`
	OvercastLevel      int64     `json:"overcastLevel"`
	HumidityPercentage int64     `json:"humidityPercentage"`
	TempFeelsLike      int64     `json:"temperatureFeelsLike"`
	TempWindChill      int64     `json:"temperatureWithWindChill"`
	Sunrise            time.Time `json:"sunriseTime"`
	Sunset             time.Time `json:"sunsetTime"`
}
