package resort

type SnowReport struct {
	ID              int64 `json:"id"`
	SeasonYear      int64 `json:"seasonYear"`
	TwelveHours     int64 `json:"twelveHours"`
	TwentyFourHours int64 `json:"twentyFourHours"`
	SeventyTwoHours int64 `json:"seventyTwoHours"`
	Week            int64 `json:"week"`
	BaseDepth       int64 `json:"baseDepth"`
	SeasonTotal     int64 `json:"seasonTotal"`
}
