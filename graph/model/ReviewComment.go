package model

type ReviewComment struct {
	ID       int64
	UserID   int64
	User     User
	Comment  string
	ReviewID int64
	Review   Review
}
