package parse

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"strconv"
)

// Parse turns a string into a slice of data.Value
func Parse(input string) ([]data.Value, error) {
	tokens := tokenize(input)
	data := make([]data.Value, len(tokens))
	for i, t := range tokens {
		d, err := t.toValue()
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	return data, nil
}

// tokenize splits a string into tokens for the parser.
func tokenize(program string) []token {

	l := &lexer{source: program, behavior: defaultBehavior}
	l.run()
	return l.tokens
}

// A token represents a string for the parser to try and parse into a value.
type token struct {
	typ  data.Type // indicates the type of value to attempt to parse the token into.
	body string    // the string to parse into a value.
}

func (t token) String() string {
	return fmt.Sprintf("%s: %s", t.typ, t.body)
}

func (t token) toValue() (val data.Value, err error) {
	switch t.typ {
	case data.Number:
		var n float64
		n, err = strconv.ParseFloat(t.body, 64)
		val = data.NewNumber(n)
	case data.String:
		val = data.NewString(t.body)
	default:
		err = fmt.Errorf("no parsing behavior defined for token type '%s'", t.typ)
	}
	return
}
