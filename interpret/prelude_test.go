package interpret

import (
	_ "embed"
	"testing"
)

//go:embed prelude.naiveconcat
var prelude string

func TestInterpretPreludeWords(t *testing.T) {
	testInterpret(t, preludeTestCases)
}

var preludeTestCases = []interpretTestCase{
	{
		"split_on",
		[]string{
			prelude,
			`"foo,bar,baz"`,
			`","`,
			"split_on say",
			`"Won't find the delimiter"`,
			`","`,
			`split_on say`,
		},
		"{\"foo\" \"bar,baz\"}\n{\"Won't find the delimiter\"}",
	},
	{
		"split_each",
		[]string{
			prelude,
			`"foo,bar,baz"`,
			`","`,
			`split_each say`,
		},
		`{"foo" "bar" "baz"}`,
	},
}
