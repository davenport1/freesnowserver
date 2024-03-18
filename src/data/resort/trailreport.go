package resort

import (
	"database/sql"
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

type TrailReportModel struct {
	Db *sql.DB
}
