package repo

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/yoteerakrit/assessment/domain"

	"gorm.io/driver/mysql"
)

type ScheduleRepoMySQL struct {
	db *gorm.DB
}

func NewMySQL() *gorm.DB {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	uri := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"keep1234",
		"tcp",
		"0.0.0.0",
		3306,
		"keeplearning",
	)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect mysql database : %s", err.Error())
	}

	return db
}

func NewSchedule(db *gorm.DB) *ScheduleRepoMySQL {
	return &ScheduleRepoMySQL{db: db}
}

func (d *ScheduleRepoMySQL) CreateSchedule(schedule *domain.Schedule) (*domain.Schedule, error) {
	return nil, nil
}
