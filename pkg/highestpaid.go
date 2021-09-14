package gbill

import (
	"bytes"
	"text/template"
)

type HighestPaid Filter

func NewHighestPaid(options []string) BillFilter {
	return &HighestPaid{"HighestPaid", options}
}

func (hp *HighestPaid) CalculateBill(bills []*Bill) []string {
	highestBill := *bills[0]
	for _, bill := range bills {
		if bill.Amount > highestBill.Amount {
			highestBill = *bill
		}
	}

	const format = `HIGEST_PAID:{{ .Name }}:{{ .PaidOn }}:{{ .Category }}:{{ .Amount }}`
	t := template.Must(template.New("line").Parse(format))

	buf := bytes.NewBufferString("")
	err := t.Execute(buf, highestBill)
	if err != nil {
		panic(err)
	}

	return []string{buf.String()}
}
