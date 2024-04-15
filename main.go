package main

import (
	"embed"
	"fmt"
	"net/http"

	"controllers"
	"templates"
	"views"

	"github.com/go-chi/chi/v5"
)

//go:embed robots.txt
var robot embed.FS

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	//fs := http.FileServer(http.FS(robot))
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(robot))
	})

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(
		views.ParseFS(templates.FS,
			"signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3008...")
	http.ListenAndServe(":3008", r)
}
