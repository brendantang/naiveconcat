package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cfg := replConfig{
		debugMode:    true,
		input:        bufio.NewReader(os.Stdin),
		initialDict:  std(),
		initialStack: &stack{},
	}
	log.Fatal(repl(cfg))
}

// A dictionary binds words to data. When a word is evaluated, it is substituted
// for the data it binds in the dictionary.
type dictionary struct {
	bindings map[string]datum
}

func (dict dictionary) get(word string) (d datum, ok bool) {
	d, ok = dict.bindings[word]
	return
}

// The stack is where data is stored while the program runs.
// The stack is "first-in, last-out"—you can only store data by "pushing" it
// onto the top, and you can only access data by "popping" it off the top.
type stack struct {
	data []datum
}

// push puts datum d on top of the stack and returns it.
func (s *stack) push(d datum) {
	s.data = append([]datum{d}, s.data...)
}

// pop removes the top item from the stack and returns them both.
func (s *stack) pop() (datum, error) {
	if len(s.data) < 1 {
		return datum{}, errors.New(emptyStackError)
	}
	var d datum
	d, s.data = s.data[0], s.data[1:]
	return d, nil
}

func (s *stack) String() (out string) {
	for _, d := range s.data {
		out = fmt.Sprintf("%v ", d.String()) + out
	}
	return
}

const emptyStackError = "tried to access the top item of the stack, but the stack is empty"

func interpretLine(line string, dict dictionary, s *stack) (newDict dictionary, err error) {
	data, err := parse(tokenize(line))
	if err != nil {
		return dict, err
	}
	newDict, err = eval(data, dict, s)
	return
}

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

func eval(data []datum, dict dictionary, s *stack) (dictionary, error) {
	for _, d := range data {

		switch d.dataType {
		case tnumber:
			// push a number on the stack
			s.push(d)
		case tword:
			// look up a word in the dictionary
			definition, ok := dict.get(d.word)
			if !ok {
				return dict, undefinedError(d)
			}
			dict, err := eval([]datum{definition}, dict, s)
			if err != nil {
				return dict, err
			}
		case tcommand:
			// execute a command
			dict, err := d.command.execute(dict, s)
			if err != nil {
				return dict, err
			}
		}
	}
	return dict, nil
}

func undefinedError(w datum) error {
	return fmt.Errorf("the word '%s' is not defined", w.word)
}
func typeError(d datum, t dataType) error {
	return fmt.Errorf("type error: expected '%s' (%s) to be type %s ", d, d.dataType, t)
}

func std() dictionary {
	return dictionary{
		bindings: map[string]datum{
			"say": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s *stack) (dictionary, error) {
						d, err := s.pop()
						if err != nil {
							return dict, err
						}
						fmt.Println(d)
						return dict, nil
					},
				},
			},
			"+": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s *stack) (dictionary, error) {
						a, err := s.pop()
						if err != nil {
							return dict, err
						}
						if a.dataType != tnumber {
							return dict, typeError(a, tnumber)
						}
						b, err := s.pop()
						if err != nil {
							return dict, err
						}
						if b.dataType != tnumber {
							return dict, typeError(b, tnumber)
						}
						s.push(datum{dataType: tnumber, number: b.number + a.number})

						return dict, nil
					},
				},
			},
			"-": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s *stack) (dictionary, error) {
						a, err := s.pop()
						if err != nil {
							return dict, err
						}
						if a.dataType != tnumber {
							return dict, typeError(a, tnumber)
						}
						b, err := s.pop()
						if err != nil {
							return dict, err
						}
						if b.dataType != tnumber {
							return dict, typeError(b, tnumber)
						}
						s.push(datum{dataType: tnumber, number: b.number - a.number})

						return dict, nil
					},
				},
			},
			"*": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s *stack) (dictionary, error) {
						a, err := s.pop()
						if err != nil {
							return dict, err
						}
						if a.dataType != tnumber {
							return dict, typeError(a, tnumber)
						}
						b, err := s.pop()
						if err != nil {
							return dict, err
						}
						if b.dataType != tnumber {
							return dict, typeError(b, tnumber)
						}
						s.push(datum{dataType: tnumber, number: a.number * b.number})

						return dict, nil
					},
				},
			},
			"/": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s *stack) (dictionary, error) {
						a, err := s.pop()
						if err != nil {
							return dict, err
						}
						if a.dataType != tnumber {
							return dict, typeError(a, tnumber)
						}
						b, err := s.pop()
						if err != nil {
							return dict, err
						}
						if b.dataType != tnumber {
							return dict, typeError(b, tnumber)
						}
						s.push(datum{dataType: tnumber, number: b.number / a.number})

						return dict, nil
					},
				},
			},
		},
	}
}
