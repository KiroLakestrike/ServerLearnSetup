package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	fmt.Println("Listening on http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}
