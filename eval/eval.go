package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

// Eval takes a slice of expressions
func Eval(values []data.Value, d *data.Dictionary, s *data.Stack) error {
	for _, v := range values {

		switch v.Type {
		case data.Number:
			// push a number on the stack
			s.Push(v)
		case data.Word:
			// look up a word in the dictionary
			definition, ok := d.Get(v.Word)
			if !ok {
				return undefinedError(v)
			}
			err := Eval([]data.Value{definition}, d, s)
			if err != nil {
				return err
			}
		case data.Proc:
			// run a procedure
			err := v.Proc.Execute(d, s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func undefinedError(w data.Value) error {
	return fmt.Errorf("the word '%s' is not defined", w.Word)
}
