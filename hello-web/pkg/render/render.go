package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Fprint(w, "Error loading template: ", tmpl)
		return
	}
	parseError := parsedTemplate.Execute(w, nil)
	if parseError != nil {
		fmt.Fprint(w, "Error parsing template: ", tmpl)
		return
	}
}
