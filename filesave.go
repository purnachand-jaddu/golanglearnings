package main

import (
	"os/signal"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/file", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.GET("/attachement", func(c echo.Context) error {
		return c.Attachment("index.html", "download.html")
	})

	e.GET("/inline", func(c echo.Context) error {
		return c.Inline("index.html", "download.html")
	})
	signal.Notify()
	e.Logger.Fatal(e.Start(":8080"))
}
