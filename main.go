package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", landing)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Landing")
}
