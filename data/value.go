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
	Str       string
	Quotation []Value
	Bool      bool
}

// Type describes the internal type of a Value.
type Type int

// A naiveconcat value is a number, string, word, procedure, quotation, or
// boolean.
const (
	Number Type = iota
	String
	Word
	Proc
	Quotation
	Boolean
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
		s = "procedure"
	case Quotation:
		s = "quotation"
	case Boolean:
		s = "boolean"
	}
	return
}

// A Procedure is an executable procedure.
type Procedure struct {
	fn func(*Dictionary, *Stack) error
}

// Execute runs a Procedure.
func (proc Procedure) Execute(d *Dictionary, s *Stack) error {
	return proc.fn(d, s)
}

func (v Value) String() (s string) {
	switch v.Type {
	case Number:
		s = strconv.FormatFloat(v.Number, 'f', -1, 64)
	case Word:
		s = v.Word
	case Proc:
		s = "PROCEDURE"
	case Quotation:
		itemStrings := make([]string, len(v.Quotation))
		for i, item := range v.Quotation {
			itemStrings[i] = item.String()
		}
		s = fmt.Sprintf("{%s}", strings.Join(itemStrings, " "))
	case String:
		s = fmt.Sprintf("\"%s\"", v.Str)
	case Boolean:
		if v.Bool {
			s = "TRUE"
		} else {
			s = "FALSE"
		}
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
		Type: String,
		Str:  s,
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
func NewProc(fn func(*Dictionary, *Stack) error) Value {
	return Value{
		Type: Proc,
		Proc: Procedure{fn},
	}
}

// NewBoolean constructs a Boolean Value from a bool.
func NewBoolean(b bool) Value {
	return Value{
		Type: Boolean,
		Bool: b,
	}
}

// A TypeErr indicates when a naiveconcat Value does not have its expected type.
type TypeErr struct {
	val      Value
	expected Type
}

// NewTypeErr returns a TypeErr.
func NewTypeErr(val Value, expected Type) error {
	return TypeErr{val, expected}
}

func (e TypeErr) Error() string {
	return fmt.Sprintf(
		"type error: expected '%s' (%s) to be type %s ",
		e.val,
		e.val.Type,
		e.expected,
	)
}
