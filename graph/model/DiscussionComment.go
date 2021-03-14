package model

type DiscussionComment struct {
	ID           int64
	UserID       int64
	User         User
	Comment      string
	DiscussionID int64
	Discussion   Discussion
}
