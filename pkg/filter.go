package gbill

type Filter struct {
	Name    string
	Options []string
}

func NewFilter(n string, o []string) Filter {
	return Filter{n, o}
}

func NewFiltersFromArgs(args []string) []Filter {
	//make sure name exits out of possibles

	fs := make([]Filter, len(args))
	options := []string{} //todo

	for _, name := range args {
		fs = append(fs, NewFilter(name, options))
	}

	return fs
}
