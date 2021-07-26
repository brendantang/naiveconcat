package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// dup pops a value, then pushes it twice.
func dup(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(val)
	s.Push(val)
	return nil
}

// drop discards the top value on the stack
func drop(d *data.Dictionary, s *data.Stack) error {
	_, err := s.Pop()
	if err != nil {
		return err
	}
	return nil
}

// lambda pops a value and pushes a procedure that evaluates to that value.
func lambda(d *data.Dictionary, s *data.Stack) error {
	head, err := s.Pop()
	if err != nil {
		return err
	}
	proc := data.NewProc(
		func(d *data.Dictionary, s *data.Stack) error {
			s.Push(head)
			return apply(d, s)
		},
	)
	s.Push(proc)
	return nil
}
