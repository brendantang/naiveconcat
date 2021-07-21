package interpret

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"strings"
	_ "testing"
)

func testingConfig(src string) Config {
	return Config{
		Prompt:       "",
		DebugMode:    true,
		Input:        bufio.NewReader(strings.NewReader(src)),
		InitialDict:  builtins.StandardDictionary(),
		InitialStack: data.NewStack(),
	}
}

func ExampleInterpret() {
	d, s := builtins.StandardDictionary(), data.NewStack()
	src := `-- From '--' to the end of a line is a comment.

		"foo" 2 "bar"   -- Literal values get pushed on the stack.

		42 2 *          -- The '*' word pops two numbers off the top of the stack,
		                -- multiplies them together, and pushes the result back on the stack.

		say	        -- The 'say' word pops the top value off the stack and prints it.
		`
	Interpret(src, d, s)
	fmt.Println(s)

	// Output:
	// 84
	// bar
	// ["foo" 2]
}
func ExampleInterpret_words() {
	d, s := builtins.StandardDictionary(), data.NewStack()

	src := `{ dup * } "square" define -- stack: []
		2 square apply            -- [4]
		square apply              -- [16]
		say                       -- []
		`
	Interpret(src, d, s)
	// Output:
	// 16
}

// Let bindings are local to their enclosing quotation(s).
func ExampleInterpret_let() {
	d, s := builtins.StandardDictionary(), data.NewStack()

	src := `3 "x" let
		x say

		-- x is 2 in the outer quotation, and 1 in the inner.
		{2 "x" let {1 "x" let x say} x say} apply apply

		x say
		`

	Interpret(src, d, s)
	// Output:
	// 3
	// 2
	// 1
	// 3
}
