package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"

// Handler for our Home page
func Home(w http.ResponseWriter, r *http.Request) {

}

// Handler for our About page
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {

	// Handlers
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Listening on http://localhost%v", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
