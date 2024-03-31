package main

import (
	"parser/parser"
	"parser/repo"
)

// Demo
// There's a comment in parser/trx.go:34
// uncomment it and transactions will start to save
func main() {

	p := parser.New(repo.New(), "https://cloudflare-eth.com")

	p.Start()

}
