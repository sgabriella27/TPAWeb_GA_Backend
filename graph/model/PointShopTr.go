package model

type PointShopTr struct {
	ID     int64
	ItemID int64
	Item   Point_Item
	UserID int64
	User   User
}
