package model

type MarketGameItem struct {
	GameItemID int64
	GameItem   GameItem
	UserID     int64
	User       User
	Price      int64
	Type       string
}
