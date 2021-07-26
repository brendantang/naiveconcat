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

// length returns the length of the quotation on top of the stack without
// popping it.
func length(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.Peek()
	if err != nil {
		return err
	}
	if quot.Type != data.Quotation {
		return data.TypeError(quot, data.Quotation)
	}
	s.Push(data.NewNumber(float64(len(quot.Quotation))))

	return nil
}

// lop pops the quotation on top of the stack, then pushes its tail, then pushes
// its head.
func lop(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.Pop()
	if err != nil {
		return err
	}
	if quot.Type != data.Quotation {
		return data.TypeError(quot, data.Quotation)
	}
	head, tail := quot.Quotation[0], quot.Quotation[1:]
	s.Push(data.NewQuotation(tail...))
	s.Push(head)

	return nil
}
