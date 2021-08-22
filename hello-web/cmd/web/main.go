package main

import (
	"fmt"
	"net/http"

	"github.com/mike-wb/udemy-go/pkg/handlers"
)

const listenPort int = 8080

// main is the main application handler
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on port %d ...\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	if err != nil {
		fmt.Printf("Error listening on port: %d\n", listenPort)
		fmt.Println("ERR: ", err)
	}
}
