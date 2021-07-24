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
			//t.Log("\nDICT", d)
			t.Log("\nSTACK\t", s)
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
	{
		"a quotation gets pushed on the stack",
		&data.Stack{},
		[]data.Value{
			data.NewQuotation(
				data.NewNumber(1),
				data.NewWord("+"),
			),
		},
		data.NewStack(
			data.NewQuotation(
				data.NewNumber(1),
				data.NewWord("+"),
			),
		),
	},
	{
		"apply a quotation to evaluate it",
		data.NewStack(
			data.NewQuotation(
				data.NewNumber(1),
				data.NewWord("+"),
			),
			data.NewNumber(100),
		),
		[]data.Value{
			data.NewWord("apply"),
		},
		data.NewStack(
			data.NewNumber(101),
		),
	},
	{
		"nested quotation application",
		data.NewStack(
			data.NewQuotation(
				data.NewQuotation(
					data.NewQuotation(
						data.NewNumber(3),
					),
				),
			),
		),
		[]data.Value{
			data.NewWord("apply"),
			data.NewWord("apply"),
		},
		data.NewStack(
			data.NewQuotation(
				data.NewNumber(3),
			),
		),
	},
	{
		"define a word that evaluates to a procedure",
		&data.Stack{},
		[]data.Value{
			data.NewQuotation(
				data.NewNumber(1),
				data.NewWord("+"),
			),
			data.NewString("increment"),
			data.NewWord("define"),
			data.NewNumber(11),
			data.NewWord("increment"),
			data.NewWord("apply"),
		},
		data.NewStack(
			data.NewNumber(12),
		),
	},
	{
		"definitions are local to their enclosing quotation",
		data.NewStack(
			data.NewQuotation(
				data.NewString("outer value"),
				data.NewString("x"),
				data.NewWord("define"),
				data.NewWord("x"),
				data.NewQuotation(
					data.NewString("inner value"),
					data.NewString("x"),
					data.NewWord("define"),
					data.NewWord("x"),
					data.NewQuotation(
						data.NewString("innermost value"),
						data.NewString("x"),
						data.NewWord("define"),
						data.NewWord("x"),
					),
				),
			),
		),
		[]data.Value{
			data.NewWord("apply"),
			data.NewWord("apply"),
			data.NewWord("apply"),
		},
		data.NewStack(
			data.NewString("innermost value"),
			data.NewString("inner value"),
			data.NewString("outer value"),
		),
	},
}
