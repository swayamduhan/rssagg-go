package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swayamduhan/rssagg-go/internal/models"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := utils.Queries.CreateUser(ctx, "Swayam Duhan")
	if err != nil {
		log.Println("Unable to create user : ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unable to create user",
		})
	}

	log.Printf("Name : %v, Created At : %v, ID : %v\n", user.Name, user.CreatedAt, user.ID)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User created!",
		"api_key" : user.ApiKey,
	})
}

func GetUserByApiKey(c echo.Context) error {
	var payload models.GetUserByApiKeyModel
	ctx := c.Request().Context()

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect body!")
	}

	user, err := utils.Queries.GetUserByApiKey(ctx, payload.ApiKey)
	if err != nil {
		return c.JSON(http.StatusNotFound, "cannot find user / internal error!")
	}

	return c.JSON(http.StatusOK, user)
}