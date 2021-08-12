package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// apply pops a value and, if it's a quotation, evaluates each of its items. If
// the value is not a quotation, apply evaluates it normally.
func apply(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}

	// evaluate each item in a quotation
	if val.Type == data.Quotation {
		// create a new dict for bindings local to this
		// quotation
		local := data.NewDictionary(d, make(map[string]data.Value))

		for _, item := range val.Quotation {
			err := Eval(item, local, s)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// non-quotation values just get evaluated
	err = Eval(val, d, s)
	if err != nil {
		return err
	}
	return nil
}
