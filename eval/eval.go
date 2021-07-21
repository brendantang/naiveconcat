package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

// Eval evaluates a value in the context of a Dictionary and Stack, possibly mutating
// both.
func Eval(val data.Value, d *data.Dictionary, s *data.Stack) error {

	switch val.Type {
	case data.Number, data.String, data.Quotation:
		// push a literal value on the stack
		s.Push(val)
	case data.Word:
		// handle special `apply` keyword
		if val.Word == "apply" {
			quot, err := s.Pop()
			if err != nil {
				return err
			}
			if quot.Type != data.Quotation {
				return data.TypeError(quot, data.Quotation)
			}

			for _, val := range quot.Quotation {
				err := Eval(val, d, s)
				if err != nil {
					return err
				}
			}

			return nil

		}
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
