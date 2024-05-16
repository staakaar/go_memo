//go:build ignore

package main

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed static/logo.png
var contents []byte

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/png", contents)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
