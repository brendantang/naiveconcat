package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

func Add(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	a, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if a.Type != data.Number {
		return dict, data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if b.Type != data.Number {
		return dict, data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number + a.Number))

	return dict, nil
}

func subtract(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	a, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if a.Type != data.Number {
		return dict, data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if b.Type != data.Number {
		return dict, data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number - a.Number))

	return dict, nil
}

func multiply(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	a, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if a.Type != data.Number {
		return dict, data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if b.Type != data.Number {
		return dict, data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(a.Number * b.Number))

	return dict, nil
}

func divide(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	a, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if a.Type != data.Number {
		return dict, data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return dict, err
	}
	if b.Type != data.Number {
		return dict, data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number / a.Number))

	return dict, nil
}
