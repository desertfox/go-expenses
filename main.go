package main

import (
	"flag"

	gbill "github.com/go-expenses/pkg"
)

type filterArgs []string

var (
	csvFile    string
	filterargs filterArgs
	//defaultPort                           string
	//loadServer, loadFromDb, mergeCsvAndDb bool
)

func (i *filterArgs) String() string {
	return ""
}

func (i *filterArgs) Set(arg string) error {
	*i = append(*i, arg)
	return nil
}

func init() {
	flag.StringVar(&csvFile, "csv", "", "CSV file to load into gBill")
	flag.Var(&filterargs, "filter", "Filters to be applied")
	//flag.BoolVar(&loadServer, "server", false, "Starts up gBill web interface")
	//flag.StringVar(&defaultPort, "p", "8080", "Port for gBill web interface")

	//flag.BoolVar(&loadFromDb, "loaddb", true, "Attempts to load local sqllite database $NAME-to-do")
	//flag.BoolVar(&mergeCsvAndDb, "merge", true, "If csv file is provided Bills parsed will be added to those from the database")
}

func main() {
	flag.Parse()

	wallet := &gbill.Wallet{}

	wallet.LoadBills(csvFile)

	wallet.LoadFilters(filterargs)

	wallet.LoadOutput()

	wallet.FlashCash()
}

/*

	s := gbill.NewFilter()

	w := gbill.NewWallet(s)

	if csvFile != "" {
		err := loadCSV(csvFile, w)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("%#v", w)

	if loadServer {
		startServer(defaultPort)
	}

}
*/
