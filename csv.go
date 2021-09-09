package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"

	gbill "github.com/go-expenses/pkg"
)

func loadCSV(csvFile string, wallet *gbill.Wallet) error {
	log.Printf("loadCSV: %v", csvFile)
	if _, err := os.Stat(csvFile); err != nil {
		log.Printf("CSV file does not exist")
		return err
	}

	dataBytes, err := ioutil.ReadFile(csvFile)
	if err != nil {
		panic(err)
	}

	line := 0
	r := csv.NewReader(bytes.NewReader(dataBytes))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line++

		if line == 1 {
			continue
		}

		bill, err := gbill.NewBill(record)
		if err != nil {
			return err
		}
		log.Printf("%#v", bill)

		wallet.AddBuild(bill)
	}

	return nil
}
