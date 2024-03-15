package resort

import (
	"freesnow/common"
	"time"
)

type TrailReport struct {
	ID          int64                  `json:"id"`
	SkiResortId int64                  `json:"skiResortId"`
	TrailName   string                 `json:"trailName"`
	Difficulty  common.TrailDifficulty `json:"difficulty"`
	Status      common.TrailStatus     `json:"status"`
	LastUpdated time.Time              `json:"lastUpdated"`
}
