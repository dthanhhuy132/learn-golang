package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders template using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates___________:", err)
		return
	}
}
