package interpret

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInterpret(t *testing.T) {
	for _, c := range interpretTestCases {

		// capture stdout during the test case
		normalStdout := os.Stdout
		tmp, err := ioutil.TempFile("", "captured_output")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmp.Name())
		os.Stdout = tmp

		// interpret the test case
		d, s := builtins.Dict(), data.NewStack()
		err = Interpret(c.src, d, s, false)
		if err != nil {
			t.Fatal(err)
		}
		output, err := ioutil.ReadFile(tmp.Name())
		if err != nil {
			t.Fatal(err)
		}
		got := strings.Trim(string(output), "\n\r \t")
		if got != c.want {
			t.Fatalf("FAIL: %s\nwant: %#v\ngot: %#v\n", c.description, c.want, got)
		}

		// set stdout back to normal
		os.Stdout = normalStdout
	}
}

var interpretTestCases = []struct {
	description string
	src         string
	want        string
}{
	{
		"some values",
		`-- From '--' to the end of a line is a comment.

		"foo" 2 "bar"   -- Literal values get pushed on the stack.

		42 2 *          -- The '*' word pops two numbers off the top of the stack,
		                -- multiplies them together, and pushes the result back on the stack.

		say	        -- The 'say' word pops the top value off the stack and prints it.
		`,
		"84",
	},
	{
		"use `then` to implement `if` with consequent and alternative",
		`
		{ 
		  "predicate" define 
		  "alternative" define
		  "consequent" define
		  consequent predicate then
		  alternative predicate not then
		  apply
		} "if" define

		{"consequent" say} {"alternative" say} false if apply`,
		`"alternative"`,
	},
	{
		"use `then` to implement a recursive function",
		`
		{
			"x" define
			x 0 = "done" define
			{ x say } done then
			{ x say x 1 - countdown apply } done not then apply
		} "countdown" define

		5 countdown apply
		`,
		`5 4 3 2 1`,
	},
}
