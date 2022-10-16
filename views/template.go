package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLtemplate *template.Template
}

func (t Template) ExecuteTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := t.HTMLtemplate.Execute(w, data)
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "There was an error executing template", http.StatusInternalServerError)
		return
	}
}

func ParseTemplate(templateName string) (Template, error) {
	tpl, err := template.ParseFiles("./templates/" + templateName + ".html")

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %v", err) // return empty template if error
	}

	return Template{
		HTMLtemplate: tpl,
	}, nil
}
