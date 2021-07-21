package interpret

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestInterpret(t *testing.T) {
	for _, c := range testCases {
		d := builtins.StandardDictionary()
		s := data.NewStack()
		err := Interpret(c.input, d, s)
		if err != nil {
			t.Fatalf(
				"FAIL: %s\nInterpreter error: %v",
				c.description,
				err,
			)
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

var testCases = []struct {
	description string
	input       string
	wantStack   *data.Stack
}{
	{
		description: "a number",
		input:       "42",
		wantStack:   data.NewStack(data.NewNumber(42)),
	},
	{
		description: "multiple numbers",
		input:       "42\r31.4\r12.11111",
		wantStack: data.NewStack(
			data.NewNumber(12.11111),
			data.NewNumber(31.4),
			data.NewNumber(42),
		),
	},
	{
		description: "arithmetic",
		input:       "12 42 + 4 - 10 * 2 /",
		wantStack: data.NewStack(
			data.NewNumber(250),
		),
	},
	{
		description: "strings",
		input:       `"I'm a string" "I am another string!"`,
		wantStack: data.NewStack(
			data.NewString("I am another string!"),
			data.NewString("I'm a string"),
		),
	},
	{
		description: "define a word that evaluates to a number",
		input:       `55 "gf-age" define gf-age`,
		wantStack: data.NewStack(
			data.NewNumber(55),
		),
	},
	{
		description: "define only saves the top item of the stack",
		input:       `32 81 55 "gf-age" define say say gf-age`,
		wantStack: data.NewStack(
			data.NewNumber(55),
		),
	},
	{
		description: "define a word that evaluates to a procedure",
		input:       `{ 1 + } "increment" define 81 increment apply`,
		wantStack: data.NewStack(
			data.NewNumber(82),
		),
	},
	{
		description: "quotation",
		input:       "{1 2 +}",
		wantStack: data.NewStack(
			data.NewQuotation(
				data.NewNumber(1),
				data.NewNumber(2),
				data.NewWord("+"),
			),
		),
	},
}
