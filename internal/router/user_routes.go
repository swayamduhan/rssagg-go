package router

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
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
	})
}