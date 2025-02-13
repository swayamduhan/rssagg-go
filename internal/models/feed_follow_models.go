package models

import "github.com/jackc/pgx/v5/pgtype"

type SubscribeFeedModel struct {
	UserID pgtype.UUID `json:"user_id"`
	FeedID pgtype.UUID `json:"feed_id"`
}

type UnsubscribeFeedModel struct {
	UserID pgtype.UUID `json:"user_id"`
	FeedID pgtype.UUID `json:"feed_id"`
}

type GetFollowedFeedsForUserModel struct {
	UserID pgtype.UUID `json:"user_id"`
}