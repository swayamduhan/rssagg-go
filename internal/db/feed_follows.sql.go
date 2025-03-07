// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFollowedFeedsForUser = `-- name: GetFollowedFeedsForUser :many
SELECT id, created_at, updated_at, user_id, feed_id FROM followed_feeds
WHERE user_id = $1
`

func (q *Queries) GetFollowedFeedsForUser(ctx context.Context, userID pgtype.UUID) ([]FollowedFeed, error) {
	rows, err := q.db.Query(ctx, getFollowedFeedsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FollowedFeed
	for rows.Next() {
		var i FollowedFeed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const subscribeFeed = `-- name: SubscribeFeed :one
INSERT INTO followed_feeds (id, user_id, feed_id)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type SubscribeFeedParams struct {
	UserID pgtype.UUID
	FeedID pgtype.UUID
}

func (q *Queries) SubscribeFeed(ctx context.Context, arg SubscribeFeedParams) (FollowedFeed, error) {
	row := q.db.QueryRow(ctx, subscribeFeed, arg.UserID, arg.FeedID)
	var i FollowedFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const unsubscribeFeed = `-- name: UnsubscribeFeed :exec
DELETE FROM followed_feeds WHERE user_id=$1 AND feed_id=$2
`

type UnsubscribeFeedParams struct {
	UserID pgtype.UUID
	FeedID pgtype.UUID
}

func (q *Queries) UnsubscribeFeed(ctx context.Context, arg UnsubscribeFeedParams) error {
	_, err := q.db.Exec(ctx, unsubscribeFeed, arg.UserID, arg.FeedID)
	return err
}
