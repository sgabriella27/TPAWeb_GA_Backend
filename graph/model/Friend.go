package model

type Friends struct {
	UserID   int64 `json:"userID"`
	FriendID int64 `json:"friendID"`
	User     *User `json:"user"`
	Friend   *User `json:"friend"`
}
