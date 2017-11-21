package main

import flag "github.com/spf13/pflag"

var (
	freq    int
	balance float64
	inPlace bool
)

func init() {
	flag.IntVarP(&freq, "freq", "f", 10, "Polling frequency in seconds")
	flag.Float64VarP(&balance, "balance", "p", 0, "If specified, calculate price of balance")
	flag.BoolVarP(&inPlace, "in-place", "o", false, "Keep ticker in place (overwrite same line)")
	flag.Parse()
}
