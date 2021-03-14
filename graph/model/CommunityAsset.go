package model

type CommunityAsset struct {
	ID      int64
	Asset   string
	Like    int64
	Dislike int64
	UserID  int64
	User    User
}
