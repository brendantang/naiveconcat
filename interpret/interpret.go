// Package interpret bridges the gap between the parser and the evaluator.
// The interpreter takes input, passes it to the parser to parse it into data
// values, and passes those values to the evaluator to evaluate them.
package interpret

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"github.com/brendantang/naiveconcat/parse"
	"log"
	"os"
)

// Interpret takes input, parses it into expressions, and then evaluates those
// expressions to mutate the dictionary and stack.
func Interpret(input string, d *data.Dictionary, s *data.Stack, Debug bool) error {
	l := parse.NewLexer(input)
	p := parse.NewParser(l.Out)
	if Debug {
		l.Debug, p.Debug = log.New(os.Stderr, "LEX:", log.LstdFlags), log.New(os.Stderr, "PARSE:", log.LstdFlags)
	}
	go l.Run()
	go p.Run()
	for more := true; more; {
		select {
		case val, ok := <-p.Out:
			if !ok {
				more = false
				break
			}
			evalErr := eval.Eval(val, d, s)
			if evalErr != nil {
				more = false
				return evalErr
			}
		case parseErr := <-p.Errs:
			if parseErr != nil {
				more = false
				return parseErr
			}
		}
	}
	return nil
}

// Config stores configuration details for the REPL.
type Config struct {
	Prompt       string           // the string that appears when waiting for input.
	Verbose      bool             // when true, the stack is printed for each REPL loop.
	Input        *bufio.Reader    // provides the source text for the REPL to interpret.
	InitialDict  *data.Dictionary // initial dictionary when the program starts.
	InitialStack *data.Stack      // initial stack of data when the program starts.
	Debug        bool             // when true, print parser and lexer debugging info.
}

// REPL (read-eval-print-loop) starts an interactive prompt.
func REPL(cfg Config) error {
	dict, s := cfg.InitialDict, cfg.InitialStack
	for true {
		// read a line from std in
		fmt.Print(cfg.Prompt)
		input, err := cfg.Input.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Error reading input: %e", err)
		}

		// interpret the line
		err = Interpret(input, dict, s, cfg.Debug)
		if err != nil {
			return err
		}

		if cfg.Verbose {
			fmt.Printf("%s\n\n", s)
		}
	}
	return nil
}
