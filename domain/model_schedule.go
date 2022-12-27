package domain

import "time"

type Schedule struct {
	ID          string    `gorm:"column:id;primaryKey" param:"id"`
	Description string    `gorm:"column:description;"`
	CreatedOn   time.Time `gorm:"column:created_on;"`
	UpdatedOn   time.Time `gorm:"column:updated_on;"`
}

func (s *Schedule) TableName() string {
	return TableNameSchedule.String()
}
