package eval

import (
	"fmt"

	"github.com/brendantang/naiveconcat/data"
)

// Eval takes a value and decides what to do with it. Literals like numbers,
// strings, quotations, and booleans are pushed on the stack. Words are
// substituted for their definition, then evaluated. Procedures are executed.
func Eval(val data.Value, d *data.Dictionary, s *data.Stack) error {
	switch val.Type {

	// literal values are pushed on the stack
	case data.Number, data.String, data.Quotation, data.Boolean:
		s.Push(val)

	// words are subsituted for their definition, then evaluated
	case data.Word:
		definition, ok := d.Get(val.Word)
		if !ok {
			return Error{val, d, s, undefinedError(val)}
		}
		err := Eval(definition, d, s)
		if err != nil {
			return Error{val, d, s, err}
		}

	// procedures are executed
	case data.Proc:
		err := val.Proc.Execute(d, s)
		if err != nil {
			return Error{val, d, s, err}
		}
	}

	return nil
}

func undefinedError(w data.Value) error {
	return fmt.Errorf("the word '%s' is not defined", w.Word)
}

// Error wraps an error with information about the value that failed to
// evaluate, the stack, and the dictionary.
type Error struct {
	Value data.Value
	Dict  *data.Dictionary
	Stack *data.Stack
	Err   error
}

func (err Error) Error() string {
	return fmt.Sprintf("error evaluating '%s': %s", err.Value, err.Err)
}
