package interpret

import (
	"testing"
)

func TestInterpretBasics(t *testing.T) {
	testInterpret(t, basicTestCases)
}

var basicTestCases = []interpretTestCase{
	{
		"some values",
		[]string{
			`-- From '--' to the end of a line is a comment.`,
			`"foo" 2 "bar"   -- Literal values get pushed on the stack`,
			`42 2 *          -- The '*' word pops two numbers off the top of the stack,
					-- multiplies them together, and pushes the result back on the stack.`,
			`say	        -- The 'say' word pops the top value off the stack and prints it.`,
		},
		"84",
	},
	{
		"use `then` to implement `if` with consequent and alternative",
		[]string{`
		{ 
		  "predicate" let 
		  "alternative" let
		  "consequent" let
		  consequent predicate then
		  alternative predicate not then
		  apply
		} "if" define

		{"consequent" say} {"alternative" say} false if`},
		`"alternative"`,
	},
	{
		"use `then` to implement a recursive function",
		[]string{`
		{
			"x" let
			x 0 = "done" let
			{ x say } done then
			{ x say x 1 - countdown } done not then apply
		} "countdown" define

		5 countdown
		`},
		"5\n4\n3\n2\n1\n0",
	},
	{
		"fibonacci", // Not tail-recursive, will run out of memory with higher numbers
		[]string{`
		{ 
			"x" let
			{0} 0 x = then
			{1} 1 x = then
			{x 1 - fib x 2 - fib +} 0 x = 1 x = or not then apply
		} "fib" define
		10 fib say
		`},
		"55",
	},
	{
		"fibonacci tail-recursive",
		[]string{`
		{
			{
				"a" let
				"b" let
				"n" let
				{a} 0 n = then
				{b} 1 n = then 
				{ n 1 -  a b +  b  fib-tail}  
					0 n =  1 n =  or not then apply
			} "fib-tail" define

			1 0 fib-tail

		} "fib" define
		75 fib say 
		`},
		"2111485077978050",
	},
	{
		"implement `each` using `then`",
		[]string{`
			{ -- Not tail recursive, could have bad performance?
				"f" let
				length "l" let
				{
					lop 
					f apply
					{f each}  length 0 = not  then apply
				} l 0 = not then apply
			} "each" define
			{1 2 3} {say} each 
			`},
		"1\n2\n3",
	},
	{
		"import source from a file",
		[]string{`
		import (
			example-import.naiveconcat
		)
		imported-from-example say
		`},
		`"Hello from imported file"`,
	},
}
