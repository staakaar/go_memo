//go:build ignore

package main

import (
	"embed"
	_ "embed"
	"io"
	"log"
	"net/http"
	"os"
	"path"

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

//go:embed static
var local embed.FS

func init() {
	fis, err := local.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fis {
		in, err := local.Open(path.Join("static", fi.Name()))
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create("embed-" + path.Base(fi.Name()))
		if err != nil {
			log.Fatal(err)
		}

		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("expected", "embed-"+path.Base(fi.Name()))
	}
}
