package gbill

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Bill struct {
	Name     string
	Category string
	Amount   int
	PaidOn   time.Time
}

var dateFormat = "02/01/2006"

func LoadBillsFromCSV(csvFile string) ([]*Bill, error) {
	if _, err := os.Stat(csvFile); err != nil {
		log.Printf("CSV file does not exist")
		return nil, err
	}

	dataBytes, err := ioutil.ReadFile(csvFile)
	if err != nil {
		panic(err)
	}

	var bills []*Bill

	line := 0
	r := csv.NewReader(bytes.NewReader(dataBytes))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line++

		if line == 1 {
			continue
		}

		bill, err := NewBill(record)
		if err != nil {
			return nil, err
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

func NewBill(record []string) (*Bill, error) {
	var amount int
	if record[2] != "" {
		decimal, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			return nil, err
		}
		amount = int(decimal * 100)
	}

	paidOn, _ := time.Parse(dateFormat, record[0])

	return &Bill{
		Name:     record[1],
		Amount:   amount,
		PaidOn:   paidOn,
		Category: "Empty",
	}, nil
}
