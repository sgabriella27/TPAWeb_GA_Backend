package model

type CommunityAssetComment struct {
	ID               int64
	UserID           int64
	User             User
	Comment          string
	CommunityAssetID int64
	CommunityAsset   CommunityAsset
}
