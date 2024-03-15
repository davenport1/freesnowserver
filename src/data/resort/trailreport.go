package resort

import "time"

type TrailReport struct {
	ID          int64     `json:"id"`
	SkiResortId int64     `json:"skiResortId"`
	TrailName   string    `json:"trailName"`
	Difficulty  int64     `json:"difficulty"`
	Status      int64     `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}
