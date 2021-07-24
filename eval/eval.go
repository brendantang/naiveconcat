package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

// Eval evaluates a value in the context of a Dictionary and Stack, possibly mutating
// both.
func Eval(val data.Value, d *data.Dictionary, s *data.Stack) error {

	//fmt.Printf("depth:%d, stack: %-*s \n", d.Depth(), d.Depth(), s)
	//fmt.Printf("depth:%d, val: %-*s \n", d.Depth(), d.Depth(), val)
	//fmt.Printf("%    *s\n", d.Depth(), val)

	switch val.Type {
	case data.Number, data.String, data.Quotation, data.Boolean:
		// push a literal value on the stack
		s.Push(val)
	case data.Word:
		switch val.Word {
		case "apply": // handle special `apply` keyword
			quot, err := s.Pop()
			if err != nil {
				return err
			}
			apply(quot, d, s)

		case "then": // handle special `then` keyword
			predicate, err := s.Pop()
			if err != nil {
				return err
			}
			if predicate.Type != data.Boolean {
				return data.TypeError(predicate, data.Boolean)
			}
			consequent, err := s.Pop()
			if err != nil {
				return err
			}
			if predicate.Bool {
				err := Eval(consequent, d, s)
				if err != nil {
					return err
				}
			}
			return nil
		case "each": // "each" keyword iterates over the items in a quotation
			iter, err := s.Pop()
			if err != nil {
				return err
			}
			if iter.Type != data.Quotation {
				return data.TypeError(iter, data.Quotation)
			}
			items, err := s.Pop()
			if err != nil {
				return err
			}
			if items.Type != data.Quotation {
				return data.TypeError(items, data.Quotation)
			}
			for _, item := range items.Quotation {
				s.Push(item)
				err := apply(iter, d, s)
				if err != nil {
					return err
				}
			}
		default: // look up a word in the dictionary
			definition, ok := d.Get(val.Word)
			if !ok {
				return undefinedError(val)
			}
			err := Eval(definition, d, s)
			if err != nil {
				return err
			}
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

func apply(quot data.Value, d *data.Dictionary, s *data.Stack) error {
	if quot.Type != data.Quotation {
		return data.TypeError(quot, data.Quotation)
	}

	// create a new dict for bindings local to this
	// quotation
	local := data.NewDictionary(d, make(map[string]data.Value))

	for _, item := range quot.Quotation {
		err := Eval(item, local, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func undefinedError(w data.Value) error {
	return fmt.Errorf("the word '%s' is not defined", w.Word)
}
