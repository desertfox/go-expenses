package gbill

import (
	"bytes"
	"text/template"
)

type LowestPaid Filter

func NewLowestPaid(options []string) BillFilter {
	return &LowestPaid{"LowestPaid", options}
}

func (lp *LowestPaid) CalculateBill(bills []*Bill) []string {
	lowestBill := *bills[0]
	for _, bill := range bills {
		if bill.Amount < lowestBill.Amount {
			lowestBill = *bill
		}
	}

	const format = `LOWEST_PAID:{{ .Name }}:{{ .PaidOn }}:{{ .Category }}:{{ .Amount }}`
	t := template.Must(template.New("line").Parse(format))

	buf := bytes.NewBufferString("")
	err := t.Execute(buf, lowestBill)
	if err != nil {
		panic(err)
	}

	return []string{buf.String()}
}
