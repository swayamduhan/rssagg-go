package models

import "github.com/jackc/pgx/v5/pgtype"

type AddFeedModel struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DeleteFeedModel struct {
	ID pgtype.UUID `json:"id"`
}

type GetFeedsToFetchModel struct {
	Limit int `json:"limit"`
}

type MarkFeedFetchedModel struct {
	ID pgtype.UUID `json:"id"`
}