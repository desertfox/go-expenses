package main

import (
	"flag"
	"log"

	gbill "github.com/go-expenses/pkg"
)

var (
	defaultPort string
	csvFile     string
	loadServer  bool
)

func init() {
	flag.BoolVar(&loadServer, "server", false, "Starts up gBill web interface")
	flag.StringVar(&defaultPort, "p", "8080", "Port for gBill web interface")

	flag.StringVar(&csvFile, "csv", "", "CSV file to load into gBill")
}

func main() {
	flag.Parse()

	w := gbill.NewWallet()

	if csvFile != "" {
		err := loadCSV(csvFile, w) //Update to return []Bills
		if err != nil {
			panic(err)
		}
	}

	log.Printf("%#v", w)

	if loadServer {
		startServer(defaultPort)
	}
}
