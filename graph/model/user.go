package model

type User struct {
	ID               int64
	AccountName      string `gorm:"uniqueIndex"`
	Password         string
	Points           int64
	ProfilePic       string
	Country          string
	RealName         string
	DisplayName      string
	CustomURL        string
	Summary          string
	Theme            string
	FrameID          int64
	Wallet           int64
	BackgroundID     int64
	BadgeID          int64
	MiniBackgroundID int64
}
