// naiveconcat is a minimal concatenative programming language, built for fun.
// Work in progress.
package main

import (
	"bufio"
	"flag"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/interpret"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	verbose := flag.Bool("verbose", false, "Print out the stack at each REPL loop")
	flag.Parse()
	cfg := interpret.Config{
		Prompt:       "> ",
		Verbose:      *verbose,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  builtins.Dict(),
		InitialStack: data.NewStack(),
	}
	filepath := flag.Arg(0)
	if filepath != "" {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		err = interpret.Interpret(string(content), builtins.Dict(), data.NewStack())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(interpret.REPL(cfg))
	}
}
