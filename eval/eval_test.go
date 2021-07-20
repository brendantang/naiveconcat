package eval

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestEval(t *testing.T) {
	for i, c := range testCases {
		d, s := builtins.StandardDictionary(), c.stack
		for _, val := range c.vals {
			Eval(val, d, s)
			//t.Log("\nDICT", d, "\nSTACK\t", s)
		}
		if c.stack.String() != c.wantStack.String() {
			failEvalTest(t, i, c, c.stack, "")
		}

	}
}
func failEvalTest(t *testing.T, i int, c testCase, got *data.Stack, msg string) {
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.wantStack, got, msg)
}

type testCase struct {
	description string
	stack       *data.Stack
	vals        []data.Value
	wantStack   *data.Stack
}

var testCases = []testCase{
	{
		"a number",
		&data.Stack{},
		[]data.Value{data.NewNumber(24)},
		data.NewStack(data.NewNumber(24)),
	},
	{
		"multiple numbers",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(24),
			data.NewNumber(2.4),
			data.NewNumber(-1000000),
		},
		data.NewStack(
			data.NewNumber(-1000000),
			data.NewNumber(2.4),
			data.NewNumber(24),
		),
	},
	{
		"numbers and strings",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(24),
			data.NewString("foo"),
			data.NewNumber(-1000000),
		},
		data.NewStack(
			data.NewNumber(-1000000),
			data.NewString("foo"),
			data.NewNumber(24),
		),
	},
	{
		"simple arithmetic",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(24),
			data.NewNumber(2),
			data.NewWord("+"),
		},
		data.NewStack(
			data.NewNumber(26),
		),
	},
	{
		"more arithmetic",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(12),
			data.NewNumber(42),
			data.NewWord("+"),
			data.NewNumber(4),
			data.NewWord("-"),
			data.NewNumber(10),
			data.NewWord("*"),
			data.NewNumber(-2.5),
			data.NewWord("/"),
		},
		data.NewStack(
			data.NewNumber(-200),
		),
	},
	{
		"defining a word",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(99),
			data.NewString("cool-num"),
			data.NewWord("define"),
			data.NewWord("cool-num"),
		},
		data.NewStack(
			data.NewNumber(99),
		),
	},
	{
		"evaluating a word and operating on it",
		&data.Stack{},
		[]data.Value{
			data.NewNumber(99),
			data.NewString("cool-num"),
			data.NewWord("define"),
			data.NewWord("cool-num"),
			data.NewNumber(1),
			data.NewWord("+"),
		},
		data.NewStack(
			data.NewNumber(100),
		),
	},
}
