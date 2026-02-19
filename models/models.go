package models

import "time"

type Task struct {
	ID        int       `db:"id"         json:"id"`
	Title     string    `db:"title"      json:"title"`
	Done      bool      `db:"done"       json:"done"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time"   json:"end_time"`
}
