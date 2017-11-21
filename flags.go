package main

import flag "github.com/spf13/pflag"

var (
	freq    int
	btc     float64
	inPlace bool
)

func init() {

	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.Float64Var(&btc, "btc", 0, "Current bitcoin balance")
	flag.BoolVar(&inPlace, "in-place", false, "Keep ticker in place by attempting to overwrite the line rather than printing on many lines (may not always work)")
	flag.Parse()

}
