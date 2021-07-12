package parse

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"strconv"
	"strings"
	"text/scanner"
)

// Parse turns a slice of tokens into a slice of data
func Parse(tokens []token) ([]data.Value, error) {
	data := make([]data.Value, len(tokens))
	for i, t := range tokens {
		d, err := t.toDatum()
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	return data, nil
}

// Tokenize splits a string into tokens for the parser.
func Tokenize(program string) (tokens []token) {
	var s scanner.Scanner
	s.Init(strings.NewReader(program))
	s.Filename = "default"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, token(s.TokenText()))
	}
	return

}

type token string

func (t token) toDatum() (data.Value, error) {
	d, err := t.toNumber()

	if err == nil {
		return d, err
	}
	// otherwise it must be a word
	d, err = t.toWord()
	return d, err
}

func (t token) toNumber() (data.Value, error) {
	n, err := strconv.ParseFloat(string(t), 64)
	if err != nil {
		err = fmt.Errorf("could not convert '%s' into a number", t)
		return data.NewNumber(0), err
	}
	d := data.NewNumber(n)
	return d, err
}

func (t token) toWord() (d data.Value, err error) {
	return data.NewWord(string(t)), nil
}

func (t token) toQuotation() (d data.Value, err error) {
	panic("implement token.toQuotation()")
}
