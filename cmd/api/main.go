package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swayamduhan/rssagg-go/internal/router"
)

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", router.TestRoutes)


	fmt.Println("Starting server ...")
	e.Logger.Fatal(e.Start(":8000"))
}