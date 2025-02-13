package models

import "github.com/jackc/pgx/v5/pgtype"

type GetPostsForUserModel struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit int32 `json:"limit"`
}