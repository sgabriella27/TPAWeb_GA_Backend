package model

import "time"

type Promo struct {
	Game_         Game `gorm:"foreignKey:GameID;constraint:OnUpdate:CASCADE,OnDelete:DO NOTHING"`
	ID            int64
	GameID        int64
	DiscountPromo int64
	EndDate       time.Time
}
