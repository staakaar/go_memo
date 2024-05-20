//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 以下と同じ http.Handle("/", http.HandleFunc(func(w http.ResponseWriter, r *http.Request)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
		switch r.Method {
		case http.MethodGet:
		default:
		}

		/** Writer */
		// contents := forecast()
		// w.Write([]byte(contents))
		f, err := os.Open("/path/to/content.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(w, f)

	})
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type MyContext struct {
	db *sql.DB
}

func (m *MyContext) handle(w http.ResponseWriter, r *http.Request) {}

func init() {
	//http.HandleFunc("/", myHandler)
	//http.ListenAndServe(":8080", nil)

	//myctx := NewMyContext()
	// http.HandleFunc("/", myctx.handle)
}
