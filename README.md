# Cryptick
A cryptocurrency ticker implemented in Golang. Get Bitcoin prices.

This project is heavily inspired by Justyn Temme's [cointick](https://github.com/justyntemme/cointick), a command line cryptocurrency ticker.

## Installation
`$ go get github.com/blunket/cryptick`

## Synopsis
```
$ ./cryptick [--freq=<seconds>] [--in-place]
$ ./cryptick [--btc=0.12345678]
```

## Usage
`$ ./cryptick [options]`

```
-p, --balance float   If specified, calculate price of balance
-f, --freq int        Polling frequency in seconds (default 10)
-o, --in-place        Keep ticker in place (overwrite same line)
```
