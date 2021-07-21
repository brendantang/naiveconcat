// naiveconcat is a minimal concatenative language.
//
//
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
	cfg := interpret.Config{
		Prompt:       "> ",
		Verbose:      *verbose,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  builtins.StandardDictionary(),
		InitialStack: data.NewStack(),
	}
	flag.Parse()
	filepath := flag.Arg(0)
	if filepath != "" {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		interpret.Interpret(string(content), builtins.StandardDictionary(), data.NewStack())
	} else {
		log.Fatal(interpret.REPL(cfg))
	}
}
