package main

import (
	"log"
	"net/http"

	api "github.com/go-expenses/pkg/server"
)

func startServer(defaultPort string) {
	http.HandleFunc("/bill", api.GetBill)
	http.HandleFunc("/bill", api.CreateBill)
	http.HandleFunc("/bill", api.UpdateBill)
	http.HandleFunc("/bill", api.DelBill)

	log.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":"+defaultPort, nil); err != nil {
		log.Fatal(err)
	}
}
