package main

import (
	"esrid/templates"
	"fmt"
	"html/template"
	"net/http"
)

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: uneven number of key-value pairs")
	}
	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}

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
		funcMap := template.FuncMap{
			"dict": dict,
		}
		t, err := template.New("").Funcs(funcMap).ParseFS(templates.TPL, "layout/layout.html", "pages/about.html", "common/*.html", "partials/*.html")
		if err != nil {
			fmt.Println("can't parse", err.Error())
		}
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
