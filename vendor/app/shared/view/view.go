package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type View struct {
	TemplateList []string
	Name         string
	TemplateDir  string
	Extention    string
	Vars         map[string]interface{}
}

func New() *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.TemplateDir = "./template/"
	v.Extention = ".gohtml"

	return v
}
func (v *View) Render(w http.ResponseWriter) {
	var tplList []string
	for _, name := range v.TemplateList {
		filePath, err := filepath.Abs(v.TemplateDir + name + v.Extention)
		if err != nil {
			log.Println(err)
		}
		tplList = append(tplList, filePath)
	}
	tc, err := template.New(v.Name).ParseFiles(tplList...)
	if err != nil {
		log.Println(err)
	}
	err = tc.ExecuteTemplate(w, v.Name+v.Extention, v.Vars)
	if err != nil {
		http.Error(w, "Template File Error : "+err.Error(), http.StatusInternalServerError)
	}

}
