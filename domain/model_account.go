package domain

type Account struct {
	ID             string `gorm:"column:id;primaryKey"`
	Name           string `gorm:"column:name;"`
	Surname        string `gorm:"column:surname;"`
	IdentityNumber string `gorm:"column:identity_number;"`
	Phone          string `gorm:"column:phone;"`
	Address        string `gorm:"column:address;"`
}
