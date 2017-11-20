package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ticker struct {
	Ask string `json:"ask"`
	Bid string `json:"bid"`
}

type coin struct {
	Ticker ticker `json:"ticker"`
}

func main() {
	var (
		freq    int     = 10
		btc     float64 = 0
		inPlace bool
	)
	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.Float64Var(&btc, "btc", 0, "Current bitcoin balance")
	flag.BoolVar(&inPlace, "in-place", false, "Keep ticker in place by attempting to overwrite the line rather than printing on many lines (may not always work)")
	flag.Parse()

	d := time.Duration(freq) * time.Second

	if inPlace {
		fmt.Println()
	}
	for {
		t := tick(btc)
		if inPlace {
			fmt.Printf("\033[1A\033[K") // move cursor up a line, then delete that line
		}
		fmt.Print(t)
		time.Sleep(d)
	}
}

func tick(btc float64) string {
	ask, bid := getTicker("btc")
	a, _ := strconv.ParseFloat(ask, 32)
	b, _ := strconv.ParseFloat(bid, 32)

	if btc != 0 {
		val := a * btc
		return fmt.Sprintf("Ask: %.2f\tBid: %.2f\tValue: %.2f\n", a, b, val)
	}

	return fmt.Sprintf("Ask: %.2f\tBid: %.2f\n", a, b)
}

func getTicker(t string) (string, string) {
	d, _ := http.Get("https://api.nexchange.io/en/api/v1/price/" + t + "USD/latest?format=json")

	c := []coin{}
	json.NewDecoder(d.Body).Decode(&c)

	ask, bid := c[0].Ticker.Ask, c[0].Ticker.Bid

	return ask, bid
}
