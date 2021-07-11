package main

import (
	"errors"
	"fmt"
)
// The stack is "first-in, last-out"—you can only store data by "pushing" it
// onto the top, and you can only access data by "popping" it off the top.
type stack struct {
	data []datum
}

// push puts datum d on top of the stack and returns it.
func (s *stack) push(d datum) {
	s.data = append([]datum{d}, s.data...)
}

// pop removes the top item from the stack and returns them both.
func (s *stack) pop() (datum, error) {
	if len(s.data) < 1 {
		return datum{}, errors.New(emptyStackError)
	}
	var d datum
	d, s.data = s.data[0], s.data[1:]
	return d, nil
}

func (s *stack) String() (out string) {
	for _, d := range s.data {
		out = fmt.Sprintf("%v ", d.String()) + out
	}
	return
}

const emptyStackError = "tried to access the top item of the stack, but the stack is empty"
