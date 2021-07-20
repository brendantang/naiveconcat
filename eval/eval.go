package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

// Eval takes a slice of expressions
func Eval(val data.Value, d *data.Dictionary, s *data.Stack) error {

	switch val.Type {
	case data.Number:
		// push a number on the stack
		s.Push(val)
	case data.Word:
		// look up a word in the dictionary
		definition, ok := d.Get(val.Word)
		if !ok {
			return undefinedError(val)
		}
		err := Eval(definition, d, s)
		if err != nil {
			return err
		}
	case data.Proc:
		// run a procedure
		err := val.Proc.Execute(d, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func undefinedError(w data.Value) error {
	return fmt.Errorf("the word '%s' is not defined", w.Word)
}
