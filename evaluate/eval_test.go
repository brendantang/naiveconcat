package evaluate

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestEval(t *testing.T) {
	for _, c := range evalTestCases {
		dict := builtins.Standard()
		s := data.NewStack()
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
	wantStack   *data.Stack
}{
	{
		description: "a number",
		inputLines: []string{
			"42",
		},
		wantStack: data.NewStack(data.NewNumber(42)),
	},
	{
		description: "multiple numbers",
		inputLines: []string{
			"42",
			"31.4",
			"12.11111",
		},
		wantStack: data.NewStack(
			data.NewNumber(12.11111),
			data.NewNumber(31.4),
			data.NewNumber(42),
		),
	},
	{
		description: "arithmetic",
		inputLines:  []string{"12 42 + 4 - 10 * 2 /"},
		wantStack: data.NewStack(
			data.NewNumber(250),
		),
	},
	{
		description: "quotation",
		inputLines:  []string{"{1 2 +}"},
		wantStack: data.NewStack(

			data.NewQuotation(
				data.NewNumber(1),
				data.NewNumber(2),
				data.NewProc(builtins.Add),
			),
		),
	},
}
