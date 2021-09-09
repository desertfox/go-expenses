package main

import (
	"fmt"
	"log"
	"net/http"

	"go-expenses/gbill/api"
)

func startServer(defaultPort string) {
	http.HandleFunc("/bill", api.GetBill)
	http.HandleFunc("/bill", api.AddBill)
	http.HandleFunc("/bill", api.UpdateBill)
	http.HandleFunc("/bill", api.DelBill)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":"+defaultPort, nil); err != nil {
		log.Fatal(err)
	}
}
