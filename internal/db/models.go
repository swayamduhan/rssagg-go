// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Feed struct {
	ID            pgtype.UUID
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	Name          string
	Url           string
	LastFetchedAt pgtype.Timestamp
}

type FollowedFeed struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	UserID    pgtype.UUID
	FeedID    pgtype.UUID
}

type Post struct {
	ID          pgtype.UUID
	FeedID      pgtype.UUID
	Title       string
	Description pgtype.Text
	Link        string
	PublishedAt pgtype.Timestamp
}

type User struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Name      string
	ApiKey    string
}
