package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("./template.html")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Mike",
		Age:  35,
		Meta: UserMeta{
			Visits: 50,
		},
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
