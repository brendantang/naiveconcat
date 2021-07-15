package data

import (
	"fmt"
	"strconv"
	"strings"
)

// A Value represents a dynamically typed piece of data.
type Value struct {
	Type      Type
	Word      string
	Number    float64
	Proc      Procedure
	String_   string
	Quotation []Value
}

// Type describes the internal type of a Value.
type Type int

const (
	Number Type = iota
	String
	Word
	Proc
	Quotation
)

func (t Type) String() (s string) {
	switch t {
	case Number:
		s = "number"
	case String:
		s = "string"
	case Word:
		s = "word"
	case Proc:
		s = "command"
	case Quotation:
		s = "quotation"
	}
	return
}

// A Procedure is an executable procedure.
type Procedure struct {
	fn func(Dictionary, *Stack) (Dictionary, error)
}

// Execute runs a Procedure.
func (proc Procedure) Execute(dict Dictionary, s *Stack) (Dictionary, error) {
	return proc.fn(dict, s)
}

func (v Value) String() (s string) {
	switch v.Type {
	case Number:
		s = strconv.FormatFloat(v.Number, 'f', -1, 64)
	case Word:
		s = v.Word
	case Proc:
		s = fmt.Sprintf("%#v", v.Proc)
	case Quotation:
		itemStrings := make([]string, len(v.Quotation))
		for i, item := range v.Quotation {
			itemStrings[i] = item.String()
		}
		s = fmt.Sprintf("{%s}", strings.Join(itemStrings, " "))
	}
	return
}

// NewNumber constructs a number Value from a float.
func NewNumber(n float64) Value {
	return Value{
		Type:   Number,
		Number: n,
	}
}

// NewString constructs a number Value from a float.
func NewString(s string) Value {
	return Value{
		Type:    String,
		String_: s,
	}
}

// NewWord constructs a word Value from a string.
func NewWord(s string) Value {
	return Value{
		Type: Word,
		Word: s,
	}
}

// NewQuotation constructs a quotation Value from other Values.
func NewQuotation(data ...Value) Value {
	return Value{
		Type:      Quotation,
		Quotation: data,
	}
}

// NewProc constructs a Proc Value from a function.
func NewProc(fn func(Dictionary, *Stack) (Dictionary, error)) Value {
	return Value{
		Type: Proc,
		Proc: Procedure{fn},
	}
}

// TypeError indicates when a Value does not have its expected type.
func TypeError(v Value, t Type) error {
	return fmt.Errorf("type error: expected '%s' (%s) to be type %s ", v, v.Type, t)
}
