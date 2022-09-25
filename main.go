package main

import (
	"fmt"
	"net/http"
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
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func printPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style={{backgroundColor: \"green\"}}>FAQ Page</h1><h2>Select a Question:</h2><select><option>kek</option><option>bur</option></select>")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/about":
		about(w, r)
	case "/path":
		printPath(w, r)
	case "/faq":
		faq(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	// http.HandleFunc("/", home)
	// http.HandleFunc("/about", about)
	// http.HandleFunc("/path", printPath)
	var router Router
	fmt.Println("Server running on :4000")
	http.ListenAndServe(":4000", router)
}
