package model

import (
	"time"

	"github.com/lib/pq"
)

type Record struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Marks     pq.Int64Array `json:"marks" gorm:"type:integer[]"`
	TotalMark int64         `json:"totalMark" gorm:"->"` // readonly (disable write permission)
	CreatedAt time.Time     `json:"createdAt"`
}

func (Record) FieldName() string {
	return "record"
}

func (Record) TableName() string {
	return "records"
}

func (Record) SelectedFields() []string {
	return []string{
		"id",
		"name",
		"marks",
		"(SELECT SUM(m) FROM UNNEST(marks) m) total_mark", // sum of marks value
		"created_at",
	}
}

type Records []Record

func (Records) FieldName() string {
	return "records"
}
