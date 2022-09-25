package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>Hello</h1>")
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<h1>Contact Page</h1>
	<a href="/">home</a>
	`+"hello")
}

func printPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style={{backgroundColor: \"green\"}}>FAQ Page</h1><h2>Select a Question:</h2><select><option>kek</option><option>bur</option></select>")
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		home(w, r)
// 	case "/about":
// 		about(w, r)
// 	case "/path":
// 		printPath(w, r)
// 	case "/faq":
// 		faq(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", home)
	router.Get("/faq", faq)
	router.Get("/about", about)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 not found", http.StatusNotFound)
	})

	fmt.Println("Server running on :4000")
	http.ListenAndServe(":4000", router)
}
