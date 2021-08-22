package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const listenPort int = 8080

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
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

// main is the main application handler
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Listening on port %d ...\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	if err != nil {
		fmt.Printf("Error listening on port: %d\n", listenPort)
		fmt.Println("ERR: ", err)
	}
}
