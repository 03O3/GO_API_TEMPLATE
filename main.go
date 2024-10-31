package main

import (
	"api/config"
	r "api/router"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	config.InitDB()

	r.SetRouters(e)

	e.Start(":1323")
}
