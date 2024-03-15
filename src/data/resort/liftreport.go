package resort

import "time"

type LiftReport struct {
	ID          int64     `json:"id"`
	SkiResortId int64     `json:"skiResortId"`
	Name        string    `json:"name"`
	Status      int64     `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}
