//go:build ignore

package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type User struct {
	Age  int
	Name string
}

type Employee struct {
	Name string
}

type Company struct {
	Employees []Employee
}

func main() {
	tmpl := `{{.}}` // HTMLをエスケープしたい時に
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, "Hello World!")
	if err != nil {
		log.Fatal(err)
	}

	user := User{Name: "Bob"}
	err = t.Execute(os.Stdout, user) // {{.Name}}

	values := []string{"Hello", "World"}
	err = t.Execute(os.Stdout, values) // {{range .}} <p>{{.}} {{index . 1}}</p> {{end}}

	user = User{Age: 21, Name: "Bob"}
	err = t.Execute(os.Stdout, user)
	/**
	 * {{ if gt .Age 20}}
	 * {{.Name}} is older than 20
	 * {{else}}
	 * {{.Name}} is not older than 20
	 * {{end}}
	 */

	company := Company{
		Employees: []Employee{
			{Name: "Bob"},
			{Name: "Mike"},
		},
	}

	err = t.Execute(os.Stdout, company)
	/// {{ with index .Employees 0}}
	/// {{.Nmae}}
	/// {{end}}

	/// {{with .}}
	// {{.}}
	/// {{else}}
	/// Not found
	/// {{end}}

	/// {{with $v := index .Employees 0}}
	/// {{$v.Name}}
	/// {{end}}

	tmpl = `<div>{{.}}</div>`
	t = template.Must(template.New("").Parse(tmpl))
	err = t.Execute(os.Stdout, template.HTML(`<b>HTML</b>`))
	// <div><b>HTML</b></div>

	tmpl = `<script>{{.}}</script>`
	t = template.Must(template.New("").Parse(tmpl))
	err = t.Execute(os.Stdout, template.JS(`alert("<script>1</script>")`))
	// <script>alert("<script>1</script>")</script>

	t = template.New("").Funcs(template.FuncMap{
		"FormatDateTime": func(fomart string, d time.Time) string {
			if d.IsZero() {
				return ""
			}
			return d.Format(fomart)
		}})

	tmpl = `{{FormatDateTime "..." .}}`
	t = template.Must(t.Parse(tmpl))
	err = t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	/// {{FomartDateTime "2006年01月02日" .}}

	t = template.Must(template.New("").ParseFiles(
		"template/index.html.tmpl",
		"template/description.html.tmpl",
		"template/login.html.tmpl",
	))

	t = template.Must(template.New("").ParseGlob("template/*.tmpl"))

	/**
	 * body.html.tmpl
	 * {{define "body"}}
	 * {{.}}
	 * {{end}}
	 *
	 * index.html.tmpl
	 * {{define "index"}}
	 * ここからが本文です
	 * {{template "body" .}} {{-template "body" .}} <= -をつけると前の行が削除される
	 * ここまでが本文です
	 * {{end}}
	 */
}
