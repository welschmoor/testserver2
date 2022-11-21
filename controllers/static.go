package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML //dangerous. This allows HTML. If users write it, dangerous
	}{
		{
			Question: "You like?",
			Answer:   "Yes",
		},
		{
			Question: "You hate?",
			Answer:   "No",
		},
		{
			Question: "Milk from the tits?",
			Answer:   "Milk from the tits!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, questions)
	}
}
