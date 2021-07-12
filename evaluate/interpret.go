package evaluate

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/parse"
)

func interpretLine(line string, dict data.Dictionary, s *data.Stack) (newDict data.Dictionary, err error) {
	data, err := parse.Parse(parse.Tokenize(line))
	if err != nil {
		return dict, err
	}
	newDict, err = eval(data, dict, s)
	return
}

func eval(values []data.Value, dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	for _, d := range values {

		switch d.Type {
		case data.Number:
			// push a number on the data.Stack
			s.Push(d)
		case data.Word:
			// look up a word in the data.Dictionary
			definition, ok := dict.Get(d.Word)
			if !ok {
				return dict, undefinedError(d)
			}
			dict, err := eval([]data.Value{definition}, dict, s)
			if err != nil {
				return dict, err
			}
		case data.Proc:
			// execute a command
			dict, err := d.Proc.Execute(dict, s)
			if err != nil {
				return dict, err
			}
		}
	}
	return dict, nil
}

func undefinedError(w data.Value) error {
	return fmt.Errorf("the word '%s' is not defined", w.Word)
}
