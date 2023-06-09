package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// if true {
	// 	http.Error(w, "Time for tea", http.StatusTeapot)
	// 	return
	// }
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
