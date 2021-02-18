package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Article struct {
	title       string
	description string
	content     string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
