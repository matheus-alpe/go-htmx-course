package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"htmxwebsite.dev/matheus-alpe/internal/templates"
)

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = templates.New()

	count := Count{Count: 0}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index", count)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
