package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

func define(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	wordName, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if wordName.Type != data.String {
		return dict, data.TypeError(wordName, data.String)
	}
	definition, err := s.Pop()
	if err != nil {
		return dict, err
	}

	newDict := dict.Set(wordName.String_, definition)

	return newDict, nil

}
