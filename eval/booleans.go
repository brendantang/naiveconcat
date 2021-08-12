package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// not pops a boolean and pushes its negation.
func not(d *data.Dictionary, s *data.Stack) error {
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(!b.Bool))
	return nil
}

// or pops two booleans and pushes TRUE if either of them is TRUE, FALSE
// otherwise.
func or(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Bool || a.Bool))
	return nil
}

// and pops two booleans and pushes TRUE if both of them are TRUE, FALSE
// otherwise.
func and(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Bool && a.Bool))
	return nil
}

// then pops a predicate and a value. If the predicate is TRUE, push the value.
// Otherwise discard it.
func then(d *data.Dictionary, s *data.Stack) error {
	predicate, err := s.PopType(data.Boolean)
	if err != nil {
		return err
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
}
