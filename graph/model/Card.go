package model

type Card struct {
	ID      int64       `json:"id"`
	BadgeID int64       `json:"badgeID"`
	Badge   *Point_Item `json:"badge"`
	//OwnedBadge []*Point_Item `json:"ownedBadge"`
	CardImg string `json:"cardImg"`
	Status  string
}
