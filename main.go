package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	bio := `<script>alert("Haha, you have been hac0rd");</script>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to Pet Search</h1><p>Bio:"+bio+"</p><p><a href=\"/contact\">Contact</a><p><a href=\"/faq\">faq</a>")

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

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", r)
}
