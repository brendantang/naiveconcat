package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

func add(d *data.Dictionary, s *data.Stack) error {
	a, err := s.Pop()
	if err != nil {
		return err
	}
	if a.Type != data.Number {
		return data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Number {
		return data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number + a.Number))

	return nil
}

func subtract(d *data.Dictionary, s *data.Stack) error {
	a, err := s.Pop()
	if err != nil {
		return err
	}
	if a.Type != data.Number {
		return data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Number {
		return data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number - a.Number))

	return nil
}

func multiply(d *data.Dictionary, s *data.Stack) error {
	a, err := s.Pop()
	if err != nil {
		return err
	}
	if a.Type != data.Number {
		return data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Number {
		return data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(a.Number * b.Number))

	return nil
}

func divide(d *data.Dictionary, s *data.Stack) error {
	a, err := s.Pop()
	if err != nil {
		return err
	}
	if a.Type != data.Number {
		return data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Number {
		return data.TypeError(b, data.Number)
	}
	s.Push(data.NewNumber(b.Number / a.Number))

	return nil
}

func equal(d *data.Dictionary, s *data.Stack) error {
	a, err := s.Pop()
	if err != nil {
		return err
	}
	if a.Type != data.Number {
		return data.TypeError(a, data.Number)
	}
	b, err := s.Pop()
	if err != nil {
		return err
	}
	if b.Type != data.Number {
		return data.TypeError(b, data.Number)
	}
	s.Push(data.NewBoolean(b.Number == a.Number))

	return nil
}
