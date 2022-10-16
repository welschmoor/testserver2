package controllers

import (
	"net/http"

	"github.com/welschmoor/testserver2/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, nil)
	}
}
