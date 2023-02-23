package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to APA Relaunch</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("startin server on :3001...")
	http.ListenAndServe(":3001", nil)
}
