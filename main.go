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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul><li><h2>Free version?</h2>30-day trial available</li><li><h2>Hours?</h2>24/7</li><li><h2>Support line?</h2>888-888-8888</li><li><h2>Monthly cost?</h2>$10</li><li><h2>Mission?</h2>Live to serve</li><li><h2>Location?</h2>Not on Earth</li></ul>
	`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Contact Us: <a href=\"mailto:paul@paul.com\">paul@paul.com</a>.")
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
