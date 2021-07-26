package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

func let(d *data.Dictionary, s *data.Stack) error {
	wordName, err := s.Pop()
	if err != nil {
		return err
	}
	if wordName.Type != data.String {
		return data.TypeError(wordName, data.String)
	}
	definition, err := s.Pop()
	if err != nil {
		return err
	}

	d.Set(wordName.Str, definition)
	return nil
}
