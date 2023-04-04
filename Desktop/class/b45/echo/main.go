package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/hello", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html")
		c.Response().WriteHeader(404)
		return c.String(405, "Halo cuy!")
	})

	fmt.Println("Server is running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
