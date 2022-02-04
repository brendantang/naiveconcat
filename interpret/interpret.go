// Package interpret bridges the gap between the parser and the evaluator.
// The interpreter takes input, passes it to the parser to parse it into data
// values, and passes those values to the evaluator to evaluate them.
package interpret

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"github.com/brendantang/naiveconcat/parse"
)

// Interpret takes input, parses it into expressions, and then evaluates those
// expressions to mutate the dictionary and stack.
func Interpret(input string, d *data.Dictionary, s *data.Stack) error {

	// First interpret any imported files in order
	if strings.HasPrefix(strings.TrimLeft(input, "\t\n\r "), "import (") {
		split := strings.SplitN(input, ")", 2)
		imports := strings.Fields(split[0])[2:] // drop the "import ("
		for _, path := range imports {
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return importErr(path, err)
			}
			err = Interpret(string(file), d, s)
			if err != nil {
				return importErr(path, err)
			}
		}
		input = split[1]
	}

	// Concurrently run the lexer and parser
	l := parse.NewLexer(input)
	p := parse.NewParser(l.Out)
	go l.Run()
	go p.Run()

	for more := true; more; {
		select {

		// Evaluate the next value from the parser
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

		// Handle a parsing error
		case parseErr := <-p.Errs:
			if parseErr != nil {
				more = false
				return parseErr
			}

		// Handle a lexing error
		case lexErr := <-l.Errs:
			if lexErr != nil {
				more = false
				return lexErr
			}

		// If both parser and lexer idle for 5 seconds, time out
		case <-time.After(5 * time.Second):
			more = false
			return timeout(l, p)
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
}

// REPL (read-eval-print-loop) starts an interactive prompt.
func REPL(cfg Config) error {
	dict, s := cfg.InitialDict, cfg.InitialStack
	for true {

		// Read a line from std in
		fmt.Print(cfg.Prompt)
		input, err := cfg.Input.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		// Interpret the line
		err = Interpret(input, dict, s)
		if err != nil {
			return err
		}

		if cfg.Verbose {
			fmt.Printf("%s\n\n", s)
		}
	}
	return nil
}

func importErr(path string, err error) error {
	return fmt.Errorf("error loading import '%s': %w", path, err)
}

func timeout(l *parse.Lexer, p *parse.Parser) error {
	return fmt.Errorf("interpreter timed out\nlexer: %s", l)
}
