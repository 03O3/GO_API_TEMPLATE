package router

import (
	"net/http"

	"api/config"
	model "api/models"

	"github.com/labstack/echo/v4"
)

func Users(c echo.Context) error {

	var users []model.User
	if err := config.DB.Table("users").Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
