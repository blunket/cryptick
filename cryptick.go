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
	var freq int = 10
	var btc float64 = 0
	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.Float64Var(&btc, "btc", 0, "Current bitcoin balance")
	flag.Parse()

	for {
		tick(freq, btc)
		time.Sleep(time.Duration(freq) * time.Second)
	}
}

func tick(freq int, btc float64) {
	ask, bid := getTicker("btc")
	a, _ := strconv.ParseFloat(ask, 32)
	b, _ := strconv.ParseFloat(bid, 32)

	if btc != 0 {
		val := a * btc
		fmt.Printf("Ask: %.2f\tBid: %.2f\tValue: %.2f\n", a, b, val)
	} else {
		fmt.Printf("Ask: %.2f\tBid: %.2f\n", a, b)
	}

}

func getTicker(t string) (string, string) {
	d, _ := http.Get("https://api.nexchange.io/en/api/v1/price/" + t + "USD/latest?format=json")

	c := []coin{}
	json.NewDecoder(d.Body).Decode(&c)

	ask, bid := c[0].Ticker.Ask, c[0].Ticker.Bid

	return ask, bid
}
