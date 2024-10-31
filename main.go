package main

import (
	"api/config"
	r "api/router"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	e := echo.New()

	config.InitDB()

	r.SetRouters(e)

	err := e.Start(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")))

	if err != nil {
		log.Fatal(err)
	}
}
