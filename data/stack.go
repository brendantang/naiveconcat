package data

import (
	"errors"
	"fmt"
)

// The Stack is "first-in, last-out"—you can only store data by "pushing" it
// onto the top, and you can only access data by "popping" it off the top.
type Stack struct {
	data []Value
}

// NewStack returns an empty Stack.
func NewStack(values ...Value) *Stack {
	return &Stack{values}
}

// Push puts Value d on top of the Stack and returns it.
func (s *Stack) Push(d Value) {
	s.data = append([]Value{d}, s.data...)
}

// Pop removes the top item from the Stack and returns them both.
func (s *Stack) Pop() (Value, error) {
	if len(s.data) < 1 {
		return Value{}, errors.New(emptyStackError)
	}
	var d Value
	d, s.data = s.data[0], s.data[1:]
	return d, nil
}

func (s *Stack) String() (out string) {
	for _, v := range s.data {
		out = fmt.Sprintf("%v ", v.String()) + out
	}
	return
}

const emptyStackError = "tried to access the top item of the Stack, but the Stack is empty"