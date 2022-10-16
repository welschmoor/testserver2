package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/welschmoor/testserver2/controllers"
	"github.com/welschmoor/testserver2/views"
)

type AboutPageInfo struct {
	CompanyName string
	AddressLine string
	Zip         int
	City        string
	LoggedIn    bool
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	tpl, err := views.ParseTemplate("home")
	if err != nil {
		panic(err)
	}
	router.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate("about")
	if err != nil {
		panic(err)
	}
	router.Get("/about", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate("faq")
	if err != nil {
		panic(err)
	}
	router.Get("/faq", controllers.StaticHandler(tpl))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 not found", http.StatusNotFound)
	})

	fmt.Println("Server running on :4000")
	http.ListenAndServe(":4000", router)
}
