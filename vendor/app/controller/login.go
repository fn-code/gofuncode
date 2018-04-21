package controller

import (
	"app/shared/view"
	"net/http"
)

type AccountLogin struct {
	Username string
	Password string
}

// Login page
func Login(w http.ResponseWriter, r *http.Request) {

	v := view.New()

	var tplList = []string{"login"}
	v.TemplateList = append(v.TemplateList, tplList...)
	v.Name = "login"

	v.Render(w)

}
