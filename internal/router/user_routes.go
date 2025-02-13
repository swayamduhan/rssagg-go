package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/controller"
)

func UserRouter(g *echo.Group) {
	userGroup := g.Group("/user")

	userGroup.GET("/create", controller.CreateUser)
	userGroup.POST("/get-by-key", controller.GetUserByApiKey)
}

