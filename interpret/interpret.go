// Package interpret bridges the gap between the parser and the evaluator.
// The interpreter takes input, passes it to the parser to parse it into data
// values, and passes those values to the evaluator to evaluate them.
package interpret

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"github.com/brendantang/naiveconcat/parse"
	"os"
)

// Interpret takes input, parses it into expressions, and then evaluates those
// expressions to mutate the dictionary and stack.
func Interpret(input string, d *data.Dictionary, s *data.Stack) error {
	vals, err := parse.Parse(input)
	if err != nil {
		return err
	}
	err = eval.Eval(vals, d, s)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	DebugMode    bool
	Input        *bufio.Reader
	InitialDict  *data.Dictionary
	InitialStack *data.Stack
}

func DefaultConfig() Config {
	return Config{
		DebugMode:    true,
		Input:        bufio.NewReader(os.Stdin),
		InitialDict:  builtins.StandardDictionary(),
		InitialStack: data.NewStack(),
	}
}

// REPL (read-eval-print-loop) starts an interactive prompt.
func REPL(cfg Config) error {
	dict, s := cfg.InitialDict, cfg.InitialStack
	for true {
		// read a line from std in
		fmt.Print("> ")
		input, err := cfg.Input.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Error reading input: %e", err)
		}

		// interpret the line
		err = Interpret(input, dict, s)
		if err != nil {
			return err
		}

		if cfg.DebugMode {
			fmt.Println(s)
		}
	}
	return nil
}
