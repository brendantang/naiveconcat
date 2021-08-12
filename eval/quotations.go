package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// length returns the length of the quotation on top of the stack without
// popping it.
func length(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.Peek()
	if err != nil {
		return err
	}
	if quot.Type != data.Quotation {
		return data.NewTypeErr(quot, data.Quotation)
	}
	s.Push(data.NewNumber(float64(len(quot.Quotation))))

	return nil
}

// lop pops the quotation on top of the stack, then pushes its tail, then pushes
// its head.
func lop(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.PopType(data.Quotation)
	if err != nil {
		return err
	}
	head, tail := quot.Quotation[0], quot.Quotation[1:]
	s.Push(data.NewQuotation(tail...))
	s.Push(head)

	return nil
}

// appendToQuot adds a value to the end of a quotation.
func appendToQuot(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.PopType(data.Quotation)
	if err != nil {
		return err
	}
	val, err := s.Pop()
	if err != nil {
		return err
	}
	vals := append(quot.Quotation, val)
	s.Push(data.NewQuotation(vals...))
	return nil
}
