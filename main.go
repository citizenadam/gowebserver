package main

import (
	"controllers"
	"fmt"
	"net/http"
	"path/filepath"
	"views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))
	// Or inline everything and skip the `tpl` variable.
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
