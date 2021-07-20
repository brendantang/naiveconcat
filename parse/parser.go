package parse

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"strconv"
)

type parser struct {
	in   chan token      // where lexed tokens are received.
	out  chan data.Value // where parsed expressions are sent.
	done chan bool       // send true when input channel closed.
	errs chan error      // where parsing errors are sent.
}

func newParser(input chan token) *parser {
	return &parser{
		in:   input,
		out:  make(chan data.Value, 2),
		done: make(chan bool, 1),
		errs: make(chan error, 1),
	}

}

func (p *parser) run() {
	for tok := range p.in {
		switch tok.typ {
		case num:
			if n, err := strconv.ParseFloat(tok.body, 64); err != nil {
				p.out <- data.NewNumber(n)
			} else {
				p.errs <- conversionError(tok, data.Number)
			}
		}
	}
	p.done <- true
}

func conversionError(tok token, typ data.Type) error {
	return fmt.Errorf("could not parse %v as %v", tok, typ)
}

// A token represents a string for the parser to try and parse into a value.
type token struct {
	typ  tokenType // indicates the type of value to attempt to parse the token into.
	body string    // the string to parse into a value.
}

type tokenType int

const (
	num tokenType = iota
	word
	str
	openQ
	closeQ
)

func (t tokenType) String() (s string) {
	switch t {
	case num:
		s = "number"
	case str:
		s = "string"
	case word:
		s = "word"
	case openQ:
		s = "open quotation"
	case closeQ:
		s = "close quotation"
	}
	return
}

func (t token) String() string {
	return fmt.Sprintf("%s: %s", t.typ, t.body)
}

func (t token) toValue() (val data.Value, err error) {
	switch t.typ {
	case num:
		var n float64
		n, err = strconv.ParseFloat(t.body, 64)
		val = data.NewNumber(n)
	case str:
		val = data.NewString(t.body)
	case word:
		val = data.NewWord(t.body)
	default:
		err = fmt.Errorf("no parsing behavior defined for token type '%s'", t.typ)
	}
	return
}
