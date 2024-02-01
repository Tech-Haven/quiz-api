package main

import (
	"quiz-api/config"
	"quiz-api/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config := config.New()
	database.ConnectDatabase(config)

	e.Logger.Fatal(e.Start(config.Listen_Addr))

}
