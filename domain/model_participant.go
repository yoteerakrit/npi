package domain

type Participant struct {
	ID        string `gorm:"column:id;primaryKey"`
	RoundId   string `gorm:"column:round_id;"`
	AccountId string `gorm:"column:account_id;"`
}
