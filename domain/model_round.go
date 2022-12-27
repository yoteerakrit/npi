package domain

import "time"

type Round struct {
	ID          string    `gorm:"column:id;primaryKey"`
	ScheduleId  string    `gorm:"column:schedule_id;"`
	Description string    `gorm:"column:description;"`
	Date        time.Time `gorm:"column:date;"`
}
