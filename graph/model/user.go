package model

type User struct {
	ID int64
	AccountName string `gorm:"uniqueIndex"`
	Password string
}
