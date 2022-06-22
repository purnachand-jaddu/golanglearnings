package main

import (
	"net/http"

	"echoframework.com/basics/chess"
	"github.com/labstack/echo/v4"
)

func main() {
	chess.Start()
}

func writeCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "Purna"
	cookie.Value = "100"
	c.SetCookie(cookie)
}

func credentialValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !areCredentialsValid() {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}
		return next(c)
	}
}

func areCredentialsValid() bool {
	return false
}
