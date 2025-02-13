package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swayamduhan/rssagg-go/internal/router"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading envs!")
	}

	port := os.Getenv("PORT")

	utils.InitDB()


	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", router.Test)

	g := e.Group("")

	router.UserRouter(g)
	router.FeedRouter(g)
	router.FeedFollowRouter(g)
	router.PostsRouter(g)

	fmt.Println("Starting server ...")
	e.Logger.Fatal(e.Start(":" + port))
}