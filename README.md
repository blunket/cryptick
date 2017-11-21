# Cryptick
A cryptocurrency ticker implemented in Golang. Get Bitcoin prices.

This project is heavily inspired by Justyn Temme's [cointick](https://github.com/justyntemme/cointick), a command line cryptocurrency ticker.

## Installation
`$ go get github.com/blunket/cryptick`

## Synopsis
`$ ./cryptick [--freq=<seconds>] [--in-place]`
`$ ./cryptick [--btc=0.12345678]`

## Usage
`$ ./cryptick [options]`

```
--freq <seconds>  Polling frequency in seconds (default 10 seconds)
--btc=<value>     If specified, calculate price of <value> btc
--in-place        Keep ticker in place (overwrite the same line)
```
