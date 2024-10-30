package main

import (
	"esrid/templates"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFS(templates.TPL, "layout/layout.html"))
		t.ExecuteTemplate(w, "layout", nil)
	})

	r.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name": "ESRID",
		}
		t := template.Must(template.ParseFS(templates.TPL, "layout/layout.html", "pages/about.html", "partials/*.html"))
		t.ExecuteTemplate(w, "layout", data)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		return
	}
}
