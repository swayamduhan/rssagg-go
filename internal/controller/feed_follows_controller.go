package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/db"
	"github.com/swayamduhan/rssagg-go/internal/models"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func SubscribeFeed(c echo.Context) error {
	var payload models.SubscribeFeedModel
	ctx := c.Request().Context()

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body")
	}

	feed_ref, err := utils.Queries.SubscribeFeed(ctx, db.SubscribeFeedParams{
		UserID: payload.UserID,
		FeedID: payload.FeedID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to subscribe to feed")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"feed_ref" : feed_ref,
	})
}

func UnsubscribeFeed(c echo.Context) error {
	var payload models.UnsubscribeFeedModel
	ctx := c.Request().Context()

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body")
	}

	err := utils.Queries.UnsubscribeFeed(ctx, db.UnsubscribeFeedParams{
		UserID: payload.UserID,
		FeedID: payload.FeedID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to unsub from feed")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
	})
}

func GetFollowedFeedsForUser(c echo.Context) error {
	var payload models.GetFollowedFeedsForUserModel
	ctx := c.Request().Context()

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body")
	}

	feed_refs, err := utils.Queries.GetFollowedFeedsForUser(ctx, payload.UserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to subscribe to feed")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"feed_refs" : feed_refs,
	})
}

