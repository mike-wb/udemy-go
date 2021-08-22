package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mike-wb/udemy-go/pkg/config"
	"github.com/mike-wb/udemy-go/pkg/handlers"
	"github.com/mike-wb/udemy-go/pkg/render"
)

const listenPort int = 8080

// main is the main application handler
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listening on port %d ...\n", listenPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	if err != nil {
		fmt.Printf("Error listening on port: %d\n", listenPort)
		fmt.Println("ERR: ", err)
	}
}
