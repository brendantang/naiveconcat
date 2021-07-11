package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	cfg := replConfig{
		debugMode:    true,
		input:        bufio.NewReader(os.Stdin),
		initialDict:  std(),
		initialStack: &stack{},
	}
	log.Fatal(repl(cfg))
}
