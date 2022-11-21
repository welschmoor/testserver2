package controllers

import (
	"net/http"

	"github.com/welschmoor/testserver2/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.ExecuteTemplate(w, nil)
	// tpl, err := views.ParseTemplateFS(templates.FS, "signup.html", "navbar.html", "tailwind.html")
	// if err != nil {
	// 	panic(err)
	// }
	// tpl.Execute(w, nil)
	// // tpl2 := StaticHandler(tpl)
	// // tpl2.Execute(w, nil)
}

// 	tpl, err := views.ParseTemplateFS(templates.FS, "signup.html", "navbar.html", "tailwind.html")
// 	if err != nil {
// 		fmt.Print("shoot")
// 	}
// 	tpl.Execute(w, nil)
// }
