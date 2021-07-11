package main

import (
	"fmt"
	"strconv"
)

// dataType describes the internal type that a  datum has.
type dataType int

const (
	tnumber dataType = iota
	tword
	tcommand
)

func (t dataType) String() (s string) {
	switch t {
	case tnumber:
		s = "number"
	case tword:
		s = "word"
	case tcommand:
		s = "command"
	}
	return
}

// A datum represents a value.
type datum struct {
	dataType dataType
	word     string
	number   float64
	command  command
}

// A command is an executable procedure.
type command struct {
	fn func(dictionary, *stack) (dictionary, error)
}

func (cmd command) execute(dict dictionary, s *stack) (dictionary, error) {
	return cmd.fn(dict, s)
}

func (d datum) String() (s string) {
	switch d.dataType {
	case tnumber:
		s = strconv.FormatFloat(d.number, 'f', -1, 64)
	case tword:
		s = d.word
	case tcommand:
		s = fmt.Sprintf("%#v", d.command)
	}
	return
}

func mkNumber(n float64) datum {
	return datum{
		dataType: tnumber,
		number:   n,
	}
}
