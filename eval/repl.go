package eval

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

type Config struct {
	DebugMode    bool
	Input        *bufio.Reader
	InitialDict  data.Dictionary
	InitialStack *data.Stack
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
		dict, err = Interpret(input, dict, s)
		if err != nil {
			return err
		}

		if cfg.DebugMode {
			fmt.Println(s)
		}
	}
	return nil
}
