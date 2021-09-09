package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	gbill "github.com/go-expenses/pkg"
)

func loadCSV(csvFile string) error {
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
	dateFormat := "02/01/2006"

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

		var amount int
		if record[3] != "" {
			decimal, err := strconv.ParseFloat(record[3], 32)
			if err != nil {
				return err
			}
			amount = int(decimal) * 100
		}

		if record[4] != "0" {
			decimal, err := strconv.ParseFloat(record[4], 32)
			if err != nil {
				return err
			}
			amount = int(decimal * 100)
		}

		paidOn, _ := time.Parse(dateFormat, record[0])

		bill := &gbill.Bill{
			Name:   record[2],
			Amount: amount,
			PaidOn: paidOn,
		}

		log.Printf("%#v", bill)
	}

	return nil
}
