//go:build ignore

package main

import (
	"fmt"
	"net/http"
	"path"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	method  string
	matcher *regexp.Regexp
	fnc     routerFunc
}

type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range rt.items {

		if v.method == r.Method && v.matcher.MatchString((r.RequestURI)) {
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)
			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}
			v.fnc(param, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	rt := router{}

	rt.GET(`^/$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	rt.GET(`^/(?P<name>\w+)$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello: %v\n", p["name"])
	})

	rt.POST(`^/api$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/json")
		fmt.Fprintln(w, `{"status": "OK"}`)
	})

	http.ListenAndServe(":8080", &rt)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/public/", http.FileServer(http.Dir("./static")))
	// publicを取り除く
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))

	fileserver := http.StripPrefix("/public/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) == ".xls" {
			w.Header().Set("Content-Type", "application/vnd.ms-excel")
		}
		fileserver.ServeHTTP(w, r)
	})

	mimemap := map[string]string{
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".ppt":  "application/vnd.ms-powerpoint",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if typ, ok := mimemap[path.Ext(r.URL.Path)]; ok {
			w.Header().Set("Content-Type", typ)
		}
		fileserver.ServeHTTP(w, r)
	})
}
