package resort

import (
	"freesnow/common"
	"time"
)

type LiftReport struct {
	ID          int64             `json:"id"`
	SkiResortId int64             `json:"skiResortId"`
	Name        string            `json:"name"`
	Status      common.LiftStatus `json:"status"`
	LastUpdated time.Time         `json:"lastUpdated"`
}
