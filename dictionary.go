package main

import (
	"fmt"
)

// A dictionary binds words to data. When a word is evaluated, it is substituted
// for the data it binds in the dictionary.
type dictionary struct {
	bindings map[string]datum
}

func (dict dictionary) get(word string) (d datum, ok bool) {
	d, ok = dict.bindings[word]
	return
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
