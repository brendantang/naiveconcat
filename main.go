package main

import (
	"bufio"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"log"
	"os"
)

func main() {
	cfg := eval.Config{
		DebugMode:    true,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  builtins.Standard(),
		InitialStack: data.NewStack(),
	}
	log.Fatal(eval.REPL(cfg))
}
