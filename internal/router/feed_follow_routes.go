package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/controller"
)

func FeedFollowRouter(g *echo.Group){
	FeedFollowGroup := g.Group("/follow-feed")

	FeedFollowGroup.POST("/sub", controller.SubscribeFeed)
	FeedFollowGroup.DELETE("/unsub", controller.UnsubscribeFeed)
	FeedFollowGroup.POST("/get-all", controller.GetFollowedFeedsForUser)
}