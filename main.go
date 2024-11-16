package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
