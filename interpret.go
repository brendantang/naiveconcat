package main

import (
	"fmt"
)

func interpretLine(line string, dict dictionary, s *stack) (newDict dictionary, err error) {
	data, err := parse(tokenize(line))
	if err != nil {
		return dict, err
	}
	newDict, err = eval(data, dict, s)
	return
}

func eval(data []datum, dict dictionary, s *stack) (dictionary, error) {
	for _, d := range data {

		switch d.dataType {
		case tnumber:
			// push a number on the stack
			s.push(d)
		case tword:
			// look up a word in the dictionary
			definition, ok := dict.get(d.word)
			if !ok {
				return dict, undefinedError(d)
			}
			dict, err := eval([]datum{definition}, dict, s)
			if err != nil {
				return dict, err
			}
		case tcommand:
			// execute a command
			dict, err := d.command.execute(dict, s)
			if err != nil {
				return dict, err
			}
		}
	}
	return dict, nil
}

func undefinedError(w datum) error {
	return fmt.Errorf("the word '%s' is not defined", w.word)
}
func typeError(d datum, t dataType) error {
	return fmt.Errorf("type error: expected '%s' (%s) to be type %s ", d, d.dataType, t)
}
