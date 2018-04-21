package controller

import (
	"app/shared/view"
	"net/http"

	"github.com/gorilla/csrf"
)

type UserAccount struct {
	Nama     string
	Username string
	Password string
}

// Register page
func Register(w http.ResponseWriter, r *http.Request) {
	v := view.New()
	var tplList = []string{"register"}
	v.TemplateList = append(v.TemplateList, tplList...)
	v.Name = "register"
	if r.Method == "GET" {
		v.Vars["csrfField"] = csrf.TemplateField(r)
	} else if r.Method == "POST" {
		// name := html.EscapeString(r.PostFormValue("nama"))
		name := r.PostFormValue("nama")
		uname := r.PostFormValue("username")
		pass := r.PostFormValue("password")

		ua := UserAccount{name, uname, pass}
		v.Vars["data"] = ua
	}
	v.Render(w)
}
