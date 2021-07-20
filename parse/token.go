package parse

import (
	"fmt"
)

// EOF is the rune that indicates the end of input
const EOF rune = -1

// A token represents a string for the parser to try and parse into a value.
type token struct {
	typ  tokenType // indicates the type of value to attempt to parse the token into.
	body string    // the string to parse into a value.
}

type tokenType int

const (
	eof tokenType = iota
	num
	word
	str
	openQ
	closeQ
)

func (t tokenType) String() (s string) {
	switch t {
	case eof:
		s = "end of file"
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
