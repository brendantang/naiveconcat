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
			t.Fatalf("FAIL: %s\nwant: %s\ngot: %s\n", c.description, c.want, got)
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
		"5\n4\n3\n2\n1\n0",
	},
	{
		"fibonacci", // Not tail-recursive, will run out of memory with higher numbers
		`
		{ 
			"x" define
			{0} 0 x = then
			{1} 1 x = then
			{x 1 - fib apply x 2 - fib apply +} 0 x = 1 x = or not then apply
		} "fib" define
		10 fib apply say
		`,
		"55",
	},
	{
		"fibonacci tail-recursive",
		`
		{
			{
				"a" define
				"b" define
				"n" define
				{a} 0 n = then
				{b} 1 n = then 
				{ n 1 -  a b +  b  fib-tail apply}  
					0 n =  1 n =  or not then apply
			} "fib-tail" define

			1 0 fib-tail apply

		} "fib" define
		75 fib apply say 
		`,
		"2111485077978050",
	},
	{
		"implement `each` using `then`",
		`
			{ -- Not tail recursive, could have bad performance
				"f" define
				length "l" define
				{
					lop 
					f apply
					{f each apply}  length 0 = not  then apply
				} l 0 = not then apply
			} "each" define
			{1 2 3} {say} each apply
			`,
		"1\n2\n3",
	},
	{
		"import source from a file",
		`
		import (
			example-import.naiveconcat
		)
		imported-from-example say
		`,
		`"Hello from imported file"`,
	},
}
