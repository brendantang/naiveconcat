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
		`"foo" "bar" say say
		1.234567 say
		{5 +} say
		{5 +} lambda say
		`,
		"\"bar\"\n\"foo\"\n1.234567\n{5 +}\nPROCEDURE",
	},
}
