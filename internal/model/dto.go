package model

import (
	"fmt"
	"time"

	"github.com/guregu/null"
)

type RecordRequest struct {
	Name  string  `json:"name"`
	Marks []int64 `json:"marks"`
}

func (r *RecordRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

func (r RecordRequest) Record() Record {
	return Record{
		Name:  r.Name,
		Marks: r.Marks,
	}
}

type RecordFilterRequest struct {
	StartDate null.String `json:"startDate"`
	EndDate   null.String `json:"endDate"`
	MinCount  null.Int    `json:"minCount"`
	MaxCount  null.Int    `json:"maxCount"`
}

func (r RecordFilterRequest) Validate() error {
	if r.StartDate.Valid {
		if _, err := time.Parse(time.DateOnly, r.StartDate.ValueOrZero()); err != nil {
			return fmt.Errorf("endDate should be formatted in YYYY-MM-DD")
		}
	}

	if r.EndDate.Valid {
		if _, err := time.Parse(time.DateOnly, r.EndDate.ValueOrZero()); err != nil {
			return fmt.Errorf("endDate should be formatted in YYYY-MM-DD")
		}
	}

	return nil
}
