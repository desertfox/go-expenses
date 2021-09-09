package gbill

import "time"

type Bill struct {
	Name      string
	Category  string
	Frequency string
	Amount    int
	DueOn     time.Time
	PaidOn    time.Time
}

type Wallet struct {
	Bills []*Bill
}
