package main

import (
	"github.com/brendantang/naiveconcat/interpret"
	"log"
)

func main() {
	cfg := interpret.DefaultConfig()
	log.Fatal(interpret.REPL(cfg))
}
