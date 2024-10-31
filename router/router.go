package router

import (
	controller "api/controllers"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {

	e.GET("/", controller.Users)

}
