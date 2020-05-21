package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	logger "github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(logger.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.POST("/auth/login", login)

	port := fmt.Sprintf(":%d", 8084)
	e.Logger.Fatal(e.Start(port))
}
