package controller

import (
	"app/shared/view"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tpl := []string{"welcome", "home"}
	if r.URL.Path != "/" {
		http.Error(w, "Page cannot be access ", http.StatusNotFound)
	} else {

		v := view.New()

		v.TemplateList = append(v.TemplateList, tpl...)
		v.Name = "home"

		v.Render(w)
	}

}
