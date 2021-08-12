package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

func add(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number + a.Number))

	return nil
}

func subtract(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number - a.Number))

	return nil
}

func multiply(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(a.Number * b.Number))

	return nil
}

func divide(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number / a.Number))

	return nil
}

func equal(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Number == a.Number))

	return nil
}
