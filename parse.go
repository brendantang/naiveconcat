package main

import (
	"fmt"
	"strconv"
	"strings"
)

// parse turns a slice of tokens into a slice of data
func parse(tokens []token) ([]datum, error) {
	data := make([]datum, len(tokens))
	for i, t := range tokens {
		d, err := t.toDatum()
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	return data, nil
}

// tokenize splits a string into tokens for the parser.
func tokenize(program string) []token {
	// split on spaces.
	rawTokens := strings.Fields(program)
	tokens := make([]token, len(rawTokens))
	for i, t := range rawTokens {
		tokens[i] = toToken(t)
	}
	return tokens

}

type token string

func toToken(s string) token {
	return token(s)
}

func (t token) toDatum() (d datum, err error) {
	d, err = t.toNumber()
	if err == nil {
		return
	}
	// otherwise it must be a word
	d, err = t.toWord()
	return
}

func (t token) toNumber() (d datum, err error) {
	n, err := strconv.ParseFloat(string(t), 64)
	if err != nil {
		err = fmt.Errorf("could not convert '%s' into a number", t)
	}
	d = datum{
		dataType: tnumber,
		number:   n,
	}
	return
}

func (t token) toWord() (d datum, err error) {
	return datum{
		dataType: tword,
		word:     string(t),
	}, nil
}
