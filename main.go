package main

import (
	"flag"
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

	if csvFile != "" {
		err := loadCSV(csvFile) //Update to return []Bills
		if err != nil {
			panic(err)
		}
	}

	if loadServer {
		startServer(defaultPort)
	}
}
