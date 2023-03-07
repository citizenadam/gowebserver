package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There Was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There Was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "img; charset=utf-8")
	fmt.Fprint(w, "<body><img src=\"/cmd/exp/img/grid.svg\" width=\"500\" height=\"500\"></body>")
}
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/svg", svgHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", r)
}
