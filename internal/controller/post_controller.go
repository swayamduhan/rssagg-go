package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/db"
	"github.com/swayamduhan/rssagg-go/internal/models"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func GetPostsForUser(c echo.Context) error {
	ctx := c.Request().Context()

	var payload models.GetPostsForUserModel

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body!")
	}

	posts, err := utils.Queries.GetPostsForUser(ctx, db.GetPostsForUserParams{
		UserID: payload.UserID,
		Limit: payload.Limit,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to fetch posts")
	}

	return c.JSON(http.StatusOK, posts)
}