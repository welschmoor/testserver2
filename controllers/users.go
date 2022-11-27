package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	fmt.Println("data.Email", data.Email)
	u.Templates.New.ExecuteTemplate(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Form submitted")
	fmt.Fprintln(w, " ")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}

	fmt.Fprint(w, "Email: ", r.FormValue("email")) //oder  r.PostForm.Get("email")
	fmt.Fprintln(w, " ")
	fmt.Fprint(w, "password: ", r.FormValue("password"))
}
