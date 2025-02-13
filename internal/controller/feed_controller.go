package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/db"
	"github.com/swayamduhan/rssagg-go/internal/models"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func AddFeed(c echo.Context) error {
	ctx := c.Request().Context()

	var payload models.AddFeedModel

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body!")
	}

	feed, err := utils.Queries.AddFeed(ctx, db.AddFeedParams{
		Name: payload.Name,
		Url: payload.Url,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to add feed")
	}

	return c.JSON(http.StatusOK, feed)
}

func DeleteFeed(c echo.Context) error {
	ctx := c.Request().Context()

	var payload models.DeleteFeedModel

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body!")
	}

	err := utils.Queries.DeleteFeed(ctx, payload.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to add feed")
	}

	return c.JSON(http.StatusOK, "Feed delete success!")
}

func GetFeeds(c echo.Context) error {
	ctx := c.Request().Context()
	feeds, err := utils.Queries.GetFeeds(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to add feed")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status" : "success",
		"feeds" : feeds,
	})
}

func GetFeedsToFetch(c echo.Context) error {
	ctx := c.Request().Context()

	var payload models.GetFeedsToFetchModel
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "wrong body")
	}

	feeds, err := utils.Queries.GetFeedsToFetch(ctx, int32(payload.Limit))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to fetch feeds")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status" : "success",
		"feeds" : feeds,
	})
}

func MarkFeedFetched(c echo.Context) error {
	ctx := c.Request().Context()

	var payload models.MarkFeedFetchedModel
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "wrong body")
	}

	feed, err := utils.Queries.MarkFeedFetched(ctx, payload.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to mark feed fetched")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status" : "success",
		"feed" : feed,
	})
}