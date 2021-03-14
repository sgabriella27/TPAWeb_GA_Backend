package model

type Discussion struct {
	ID          int64
	User        User
	UserID      int64
	Game        Game
	GameID      int64
	Title       string
	Description string
}
