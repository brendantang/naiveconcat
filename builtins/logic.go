package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

func not(d *data.Dictionary, s *data.Stack) error {
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Boolean {
		return data.TypeError(b, data.Boolean)
	}
	s.Push(data.NewBoolean(!b.Bool))
	return nil
}
