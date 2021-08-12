package data

import (
	"fmt"
	"strings"
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

// Push puts Value d on top of the Stack.
func (s *Stack) Push(d Value) {
	s.data = append([]Value{d}, s.data...)
}

// Pop removes the top item from the Stack and returns it.
func (s *Stack) Pop() (Value, error) {
	if len(s.data) < 1 {
		return Value{}, EmptyStackErr{}
	}
	var d Value
	d, s.data = s.data[0], s.data[1:]
	return d, nil
}

// PopType takes a Type and Pops the top value. Error reports if the stack is
// empty or if the value doesn't match the given type.
func (s *Stack) PopType(t Type) (Value, error) {
	val, err := s.Pop()
	if err != nil {
		return val, err
	}
	if val.Type != t {
		return val, NewTypeErr(val, t)
	}
	return val, nil
}

// Peek returns the top item from the Stack without consuming it.
func (s *Stack) Peek() (Value, error) {
	if len(s.data) < 1 {
		return Value{}, EmptyStackErr{}
	}
	return s.data[0], nil
}

func (s *Stack) String() string {
	var strs []string
	for _, val := range s.data {
		strs = append([]string{val.String()}, strs...)
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, " "))
}

// An EmptyStackErr indicates when the user tries to access an item from the
// stack, but the stack is empty.
type EmptyStackErr struct{}

func (err EmptyStackErr) Error() string {
	return "tried to access the top item of the Stack, but the Stack is empty"
}
