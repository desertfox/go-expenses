package gbill

import (
	"strconv"
	"time"
)

var dateFormat = "02/01/2006"

func NewBill(record []string) (*Bill, error) {
	var amount int
	if record[3] != "" {
		decimal, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			return nil, err
		}
		amount = int(decimal) * 100
	}

	if record[4] != "0" {
		decimal, err := strconv.ParseFloat(record[4], 32)
		if err != nil {
			return nil, err
		}
		amount = int(decimal * 100)
	}

	paidOn, _ := time.Parse(dateFormat, record[0])

	return &Bill{
		Name:   record[2],
		Amount: amount,
		PaidOn: paidOn,
	}, nil
}
