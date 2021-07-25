package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)


func dup(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(val)
	s.Push(val)
	return nil
}
