package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// let pops a string and the next value, then saves a word named for that string
// which evaluates to that value.
func let(d *data.Dictionary, s *data.Stack) error {
	wordName, err := s.PopType(data.String)
	if err != nil {
		return err
	}
	definition, err := s.Pop()
	if err != nil {
		return err
	}

	d.Set(wordName.Str, definition)
	return nil
}

// define pops a string and the next value, then saves a word named for that string which evaluates to a procedure that applies that value.
func define(d *data.Dictionary, s *data.Stack) error {

	wordName, err := s.PopType(data.String)
	if err != nil {
		return err
	}
	definition, err := s.Pop()
	if err != nil {
		return err
	}
	proc := data.NewProc(
		func(d *data.Dictionary, s *data.Stack) error {
			s.Push(definition)
			return apply(d, s)
		},
	)
	d.Set(wordName.Str, proc)

	return nil
}
