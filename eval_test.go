package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	for _, c := range evalTestCases {
		dict := std()
		s := &stack{}
		for i, line := range c.inputLines {
			newDict, err := interpretLine(line, dict, s)
			if err != nil {
				t.Fatalf(
					"FAIL: %s\nError on input line %d: %v",
					c.description,
					i,
					err,
				)
			}
			dict = newDict
		}
		if s.String() != c.wantStack.String() {
			t.Fatalf(
				"FAIL: %s\nWant: %s\nHave: %s\n",
				c.description,
				c.wantStack,
				s,
			)

		}

	}
}

var evalTestCases = []struct {
	description string
	inputLines  []string
	wantStack   *stack
}{
	{
		description: "a number",
		inputLines: []string{
			"42",
		},
		wantStack: &stack{
			data: []datum{mkNumber(42)},
		},
	},
	{
		description: "multiple numbers",
		inputLines: []string{
			"42",
			"31.4",
			"12.11111",
		},
		wantStack: &stack{
			data: []datum{
				mkNumber(12.11111),
				mkNumber(31.4),
				mkNumber(42),
			},
		},
	},
	{
		description: "arithmetic",
		inputLines:  []string{"12 42 + 4 - 10 * 2 /"},
		wantStack: &stack{
			data: []datum{
				mkNumber(250),
			},
		},
	},
}
