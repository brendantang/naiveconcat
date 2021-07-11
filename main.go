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
	inputReader := bufio.NewReader(os.Stdin)
	log.Fatal(interactive(inputReader, std(), stack{}))
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
	fn func(dictionary, stack) (dictionary, stack, error)
}

func (cmd command) execute(dict dictionary, s stack) (dictionary, stack, error) {
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

// The stack is where data is stored while the program runs.
// The stack is "first-in, last-out"—you can only store data by "pushing" it
// onto the top, and you can only access data by "popping" it off the top.
type stack struct {
	data []datum
}

// push puts datum d on top of the stack and returns it.
func (s stack) push(d datum) stack {
	return stack{append(s.data, d)}
}

// pop removes the top item from the stack and returns them both.
func (s stack) pop() (datum, stack, error) {
	if len(s.data) < 1 {
		return datum{}, s, errors.New(emptyStackError)
	}
	d, rest := s.data[0], s.data[1:]
	s = stack{data: rest}
	return d, s, nil
}

func (s stack) String() (out string) {
	for _, d := range s.data {
		out = out + fmt.Sprintf("%v ", d.String())
	}
	return
}

const emptyStackError = "tried to access the top item of the stack, but the stack is empty"

// repl (read-eval-print-loop) starts an interactive prompt.
func interactive(in *bufio.Reader, dict dictionary, s stack) error {

	for true {
		// read a line from std in
		fmt.Print("> ")
		input, err := in.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Error reading input: %e", err)
		}

		// interpret the line
		dict, s, err = interpretLine(input, dict, s)
		if err != nil {
			return err
		}
		fmt.Println(s)
	}
	return nil
}

func interpretLine(line string, dict dictionary, s stack) (newDict dictionary, newStack stack, err error) {
	data, err := parse(tokenize(line))
	if err != nil {
		return dict, s, err
	}
	newDict, newStack, err = eval(data, dict, s)
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

func eval(data []datum, dict dictionary, s stack) (dictionary, stack, error) {
	for _, d := range data {
		switch d.dataType {
		case tnumber:
			// push a number on the stack
			s = s.push(d)
		case tword:
			// look up a word in the dictionary
			val, ok := dict.get(d.word)
			if !ok {
				return dict, s, undefinedError(d)
			}
			return eval([]datum{val}, dict, s)
		case tcommand:
			// execute a command
			dict, newS, err := d.command.execute(dict, s)
			if err != nil {
				return dict, s, err
			}
			s = newS

		}
	}
	return dict, s, nil
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
					fn: func(dict dictionary, s stack) (dictionary, stack, error) {
						d, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						fmt.Println(d)
						return dict, s, nil
					},
				},
			},
			"+": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s stack) (dictionary, stack, error) {
						a, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if a.dataType != tnumber {
							return dict, s, typeError(a, tnumber)
						}
						b, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if b.dataType != tnumber {
							return dict, s, typeError(b, tnumber)
						}
						s = s.push(datum{dataType: tnumber, number: a.number + b.number})

						return dict, s, nil
					},
				},
			},
			"-": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s stack) (dictionary, stack, error) {
						a, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if a.dataType != tnumber {
							return dict, s, typeError(a, tnumber)
						}
						b, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if b.dataType != tnumber {
							return dict, s, typeError(b, tnumber)
						}
						s = s.push(datum{dataType: tnumber, number: a.number - b.number})

						return dict, s, nil
					},
				},
			},
			"*": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s stack) (dictionary, stack, error) {
						a, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if a.dataType != tnumber {
							return dict, s, typeError(a, tnumber)
						}
						b, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if b.dataType != tnumber {
							return dict, s, typeError(b, tnumber)
						}
						s = s.push(datum{dataType: tnumber, number: a.number * b.number})

						return dict, s, nil
					},
				},
			},
			"/": {
				dataType: tcommand,
				command: command{
					fn: func(dict dictionary, s stack) (dictionary, stack, error) {
						a, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if a.dataType != tnumber {
							return dict, s, typeError(a, tnumber)
						}
						b, s, err := s.pop()
						if err != nil {
							return dict, s, err
						}
						if b.dataType != tnumber {
							return dict, s, typeError(b, tnumber)
						}
						s = s.push(datum{dataType: tnumber, number: a.number / b.number})

						return dict, s, nil
					},
				},
			},
		},
	}
}
