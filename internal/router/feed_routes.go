package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/controller"
)

func FeedRouter(g *echo.Group) {
	feedGroup := g.Group("/feed")

	feedGroup.POST("/add-feed", controller.AddFeed)
	feedGroup.DELETE("/delete", controller.DeleteFeed)
	feedGroup.GET("/all", controller.GetFeeds)
	feedGroup.POST("/fetch", controller.GetFeedsToFetch)
	feedGroup.PUT("/mark-fetched", controller.MarkFeedFetched)
}