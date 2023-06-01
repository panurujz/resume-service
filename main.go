package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panurujz/resume-service/handlers"
)

func main() {

	// _ = config.Open()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	u := e.Group("/user")
	u.POST("/save", handlers.CreateUser)

	e.Logger.Fatal(e.Start(":3001"))
}
