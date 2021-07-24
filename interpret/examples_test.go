package interpret

import (
	"fmt"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
)

func ExampleInterpret() {
	d, s := builtins.Dict(), data.NewStack()
	src := `-- From '--' to the end of a line is a comment.

		"foo" 2 "bar"   -- Literal values get pushed on the stack.

		42 2 *          -- The '*' word pops two numbers off the top of the stack,
		                -- multiplies them together, and pushes the result back on the stack.

		say	        -- The 'say' word pops the top value off the stack and prints it.
		`
	err := Interpret(src, d, s, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	// Output:
	// 84
	// ["foo" 2 "bar"]
}

func ExampleInterpret_words() {
	d, s := builtins.Dict(), data.NewStack()

	src := `{ dup * } "square" define -- stack: []
		2 square apply            -- [4]
		square apply              -- [16]
		say                       -- []
		`
	err := Interpret(src, d, s, false)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 16
}

// "Define" bindings are local to their enclosing quotation(s).
func ExampleInterpret_locals() {
	d, s := builtins.Dict(), data.NewStack()

	src := `3 "x" define
		x say

		-- x is 2 in the outer quotation, and 1 in the inner.
		{2 "x" define {1 "x" define x say} x say} apply apply

		x say
		`

	err := Interpret(src, d, s, false)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 3
	// 2
	// 1
	// 3
}

// Flow control using "then"
func ExampleInterpret_conditions() {
	d, s := builtins.Dict(), data.NewStack()

	src := `"You won't see this message" false then
		"You will see this message" true then
		say
		`

	err := Interpret(src, d, s, false)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// "You will see this message"
}
