package domain

type History struct {
	ID        string `gorm:"column:id;primaryKey"`
	RoundId   string `gorm:"column:round_id;"`
	AccountId string `gorm:"column:account_id;"`
	Score     string `gorm:"column:score;"`
	Note      string `gorm:"column:note;"`
}
