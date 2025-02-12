-- name: SubscribeFeed :one
INSERT INTO followed_feeds (id, user_id, feed_id)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: UnsubscribeFeed :exec
DELETE FROM followed_feeds WHERE user_id=$1 AND feed_id=$2;

-- name: GetFollowedFeedsForUser :many
SELECT * FROM followed_feeds
WHERE user_id = $1;