package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/controller"
)

func PostsRouter(g *echo.Group) {
	PostsGroup := g.Group("/posts")

	PostsGroup.POST("/get", controller.GetPostsForUser)
}