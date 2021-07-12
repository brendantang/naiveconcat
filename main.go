package main

import (
	"bufio"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/evaluate"
	"log"
	"os"
)

func main() {
	cfg := evaluate.Config{
		DebugMode:    true,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  builtins.Standard(),
		InitialStack: data.NewStack(),
	}
	log.Fatal(evaluate.REPL(cfg))
}
