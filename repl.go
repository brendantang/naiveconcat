package main

import (
	"bufio"
	"fmt"
)

type replConfig struct {
	debugMode    bool
	input        *bufio.Reader
	initialDict  dictionary
	initialStack *stack
}

// repl (read-eval-print-loop) starts an interactive prompt.
func repl(cfg replConfig) error {
	dict, s := cfg.initialDict, cfg.initialStack
	for true {
		// read a line from std in
		fmt.Print("> ")
		input, err := cfg.input.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Error reading input: %e", err)
		}

		// interpret the line
		dict, err = interpretLine(input, dict, s)
		if err != nil {
			return err
		}

		if cfg.debugMode {
			fmt.Println(s)
		}
	}
	return nil
}
