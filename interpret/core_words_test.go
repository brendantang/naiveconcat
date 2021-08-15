package interpret

import (
	"testing"
)

func TestInterpretCoreWords(t *testing.T) {
	testInterpret(t, coreWordTestCases)
}

var coreWordTestCases = []interpretTestCase{
	{
		"say",
		[]string{`"foo" "bar" say say
		1.234567 say
		{5 +} say
		{5 +} lambda say
		`},
		"\"bar\"\n\"foo\"\n1.234567\n{5 +}\nPROCEDURE",
	},
	{
		"stack",
		[]string{`"foo" "bar" 1.234567 {5 +} {5 +} lambda
		stack
		`},
		"[\"foo\" \"bar\" 1.234567 {5 +} PROCEDURE]",
	},
	{
		"words",
		[]string{`
		{ 3 "three" let words} apply
		`},
		`three	3`,
	},
	{
		"words with a local dict",
		[]string{`{"foo" "x" let words} apply`},
		`x	"foo"`,
	},
	{
		"append to a quotation",
		[]string{`{1 2 3} 4 append say`},
		`{1 2 3 4}`,
	},
}
