package resort

import (
	"database/sql"
	"freesnow/common"
	"time"
)

// LiftReport a struct representing all the different aspects of the lift with a foreign key to a ski resort
type LiftReport struct {
	ID          int64             `json:"id"`
	SkiResortId int64             `json:"skiResortId"`
	Name        string            `json:"name"`
	Status      common.LiftStatus `json:"status"`
	LastUpdated time.Time         `json:"lastUpdated"`
}

// LiftReportModel represents a model for managing lift reports in the database for a ski resort.
type LiftReportModel struct {
	Db *sql.DB // Db represents the database connection.
}

// SaveLiftsForNewResort saves a batch of lift reports for a new ski resort into the database.
// It takes a slice of LiftReport structs as input and iterates through each report,
// calling SaveLiftReport method for each report. If any error occurs during the process,
// it immediately returns the error.
func (l *LiftReportModel) SaveLiftsForNewResort(reports []LiftReport) error {
	for _, report := range reports {
		if err := l.SaveLiftReport(&report); err != nil {
			return err
		}
	}
	return nil
}

// UpdateAllLiftsForExistingResort updates all lift reports for an existing ski resort in the database.
// It takes a slice of LiftReport structs as input and iterates through each report,
// calling UpdateLiftReport method for each report. If any error occurs during the process,
// it immediately returns the error.
func (l *LiftReportModel) UpdateAllLiftsForExistingResort(reports []LiftReport) error {
	for _, report := range reports {
		if err := l.UpdateLiftReport(&report); err != nil {
			return err
		}
	}
	return nil
}

// SaveLiftReport saves a single lift report into the database.
// It takes a pointer to a LiftReport struct as input.
// Implementation of this method is required to interact with the database and store the lift report.
func (l *LiftReportModel) SaveLiftReport(lr *LiftReport) error {
	return nil
}

// UpdateLiftReport updates a single lift report in the database.
// It takes a pointer to a LiftReport struct as input.
// Implementation of this method is required to interact with the database and update the lift report.
func (l *LiftReportModel) UpdateLiftReport(lr *LiftReport) error {
	return nil
}
