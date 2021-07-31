// naiveconcat is a minimal concatenative programming language, built for fun.
// Work in progress.
package main

import (
	"bufio"
	_ "embed"
	"flag"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"github.com/brendantang/naiveconcat/interpret"
	"io/ioutil"
	"log"
	"os"
)

//go:embed prelude.naiveconcat
var prelude string

func main() {
	verbose := flag.Bool("verbose", false, "Print out the stack at each REPL loop")
	coreOnly := flag.Bool("core", false, "Skip the standard library, start with built-ins only")
	flag.Parse()
	dict := eval.CoreDict()
	stack := data.NewStack()
	if !*coreOnly {
		err := interpret.Interpret(prelude, dict, stack)
		if err != nil {
			log.Fatalf("error interpreting the standard library: %v", err)
		}
	}
	cfg := interpret.Config{
		Prompt:       "> ",
		Verbose:      *verbose,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  dict,
		InitialStack: stack,
	}
	filepath := flag.Arg(0)
	if filepath != "" {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		err = interpret.Interpret(string(content), dict, stack)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(interpret.REPL(cfg))
	}
}
