package model

type Review struct {
	ID          int64
	User        User
	UserID      int64
	Game        Game
	GameID      int64
	Description string
	Recommended bool
	Upvote      int64
	Downvote    int64
	Helpful     int64
	NotHelpful  int64
}
