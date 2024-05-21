//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocraft/web"
	"github.com/julienschmidt/httprouter"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, strings.Repeat("Hello", c.HelloCount), "World")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s !\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))

	router = web.New(AppContext{}).Middleware(web.LoggerMiddleware).Middleware(web.ShowErrorsMiddleware).Middleware((*AppContext).SetHelloCount).Get("/", (*AppContext).SayHello)
	http.ListenAndServe("localhost:3000", router)
}
