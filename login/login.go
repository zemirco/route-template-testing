package login

import (
	"html/template"
	"net/http"
)

var view = template.Must(template.ParseFiles(
	"login/login.html",
	"views/layout.html",
))

func GetLogin(w http.ResponseWriter, r *http.Request) {
	err := view.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		panic(err)
	}
}
