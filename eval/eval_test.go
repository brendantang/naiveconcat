package eval

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestEval(t *testing.T) {
	for i, c := range testCases {
		for _, val := range c.vals {
			Eval(val, builtins.StandardDictionary(), &c.stack)
		}
		if c.stack.String() != c.wantStackStr {
			failEvalTest(t, i, c, &c.stack, "")
		}

	}
}
func failEvalTest(t *testing.T, i int, c testCase, got *data.Stack, msg string) {
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.wantStackStr, got, msg)
}

type testCase struct {
	description  string
	stack        data.Stack
	vals         []data.Value
	wantStackStr string
}

var testCases = []testCase{
	{
		"a number",
		data.Stack{},
		[]data.Value{data.NewNumber(24)},
		"[24]",
	},
	{
		"multiple numbers",
		data.Stack{},
		[]data.Value{
			data.NewNumber(24),
			data.NewNumber(2.4),
			data.NewNumber(-1000000),
		},
		"[24 2.4 -1000000]",
	},
}
