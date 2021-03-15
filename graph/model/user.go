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
	Suspended        bool
	Reported         int64
	FrameID          int64
	Wallet           int64
	BackgroundID     int64
	BadgeID          int64
	MiniBackgroundID int64
	Level            int64
	Status           string
	FriendCode       string
	CountryID        int64
}

type UnsuspensionRequest struct {
	UserID int
	Reason string
	Status string
}
