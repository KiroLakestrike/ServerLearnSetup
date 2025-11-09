package main

import (
	"fmt"
	"net/http"

	"github.com/KiroLakestrike/bedAndBreakfast/pkg/handlers"
)

const PortNumber = ":8080"

func main() {

	// Handlers
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Listening on http://localhost%v", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
