package model

import "time"

type Game struct {
	ID int64
	GameTitle string
	GameDescription string
	GamePrice int
	CreatedAt time.Time
	GamePublisher string
	GameDeveloper string
	GameTag string
	GameSystemRequirement string
	GameAdult bool
	GameBannerID int64
	GameGameBanner GameMedia `gorm:"foreignKey:GameBannerID"`
	GameGameSlideshow []GameSlideshow `gorm:"foreignKey:GameID"`
}

type GameMedia struct {
	ID int64
	ImageVideo []byte
	Type string
}

type GameSlideshow struct {
	ID int64
	GameMediaID int64
	GameID int64
	GameSlideshowMedia GameMedia `gorm:"foreignKey:GameMediaID"`
	GameGameID Game `gorm:"foreignKey:GameID"`
}