package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to Pet Search</h1><p><a href=\"/contact\">Contact</a><p><a href=\"/faq\">faq</a>")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>
			<h2>Is there a free version</h2>
			Yes We have a 30 day trial!
		</li>
		<li>
			<h2> What are your hours</h2>
			Our hours are 24/7!
		</li>
		<li>
			<h2>What is the support line?</h2>
			The support line is 888-888-8888
		</li>
		<li>
			<h2>What is the montly cost?</h2>
			The montly cost is $10!
		</li>
		<li>
			<h2>What is your company mission?</h2>
			We live to serve.
		</li>
		<li>
			<h2>Where is this site?</h2>
			Not on Earth
		</li>
		</ul>	
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
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3001", r)
}
