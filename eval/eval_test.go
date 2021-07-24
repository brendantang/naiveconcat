package eval

import (
	"github.com/brendantang/naiveconcat/builtins"
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestEval(t *testing.T) {
	for i, c := range testCases {
		d, s := builtins.Dict(), c.stack
		for _, val := range c.vals {
			//t.Log("\nVAL", val)
			//t.Log("\nDICT", d)
			//t.Log("\nSTACK\t", s)
			err := Eval(val, d, s)
			if err != nil {
				failEvalTest(t, i, c, c.stack, err.Error())
			}

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
	{
		"booleans",
		&data.Stack{},
		[]data.Value{data.NewBoolean(false), data.NewBoolean(true)},
		data.NewStack(data.NewBoolean(true), data.NewBoolean(false)),
	},
	{
		"conditional flow control using `then`",
		data.NewStack(
			data.NewBoolean(false),
			data.NewString("This value won't be on the stack"),
			data.NewBoolean(true),
			data.NewString("Hello!"),
		),
		[]data.Value{data.NewWord("then"), data.NewWord("then")},
		data.NewStack(data.NewString("Hello!")),
	},
	{
		"`then` implements `if` with consequent and alternative",
		data.NewStack(
			data.NewQuotation(
				data.NewString("predicate"),
				data.NewWord("define"),
				data.NewString("alternative"),
				data.NewWord("define"),
				data.NewString("consequent"),
				data.NewWord("define"),
				data.NewWord("consequent"), data.NewWord("predicate"), data.NewWord("then"),
				data.NewWord("alternative"), data.NewWord("predicate"), data.NewWord("not"), data.NewWord("then"),
			),
		),
		[]data.Value{
			data.NewString("if"), data.NewWord("define"),
			data.NewString("This value will be on the stack."),
			data.NewString("This value won't."),
			data.NewBoolean(true),
			data.NewWord("if"),
			data.NewWord("apply"),
		},
		data.NewStack(data.NewString("This value will be on the stack.")),
	},
	{
		"simple recursive definition",
		data.NewStack(
			data.NewQuotation(
				data.NewString("x"),
				data.NewWord("define"),
				data.NewWord("x"), data.NewNumber(0), data.NewWord("="), data.NewString("done"), data.NewWord("define"),
				data.NewQuotation(
					data.NewWord("x"), //data.NewWord("say"),
				),
				data.NewWord("done"), data.NewWord("then"),

				data.NewQuotation(
					data.NewWord("x"), // Push on the stack
					data.NewWord("x"), data.NewNumber(1), data.NewWord("-"), data.NewWord("countdown"), data.NewWord("apply"),
				),
				data.NewWord("done"), data.NewWord("not"), data.NewWord("then"), data.NewWord("apply"),
			),
		),
		[]data.Value{
			data.NewString("countdown"), data.NewWord("define"),
			data.NewNumber(10),
			data.NewWord("countdown"),
			data.NewWord("apply"),
		},
		data.NewStack(
			data.NewNumber(0),
			data.NewNumber(1),
			data.NewNumber(2),
			data.NewNumber(3),
			data.NewNumber(4),
			data.NewNumber(5),
			data.NewNumber(6),
			data.NewNumber(7),
			data.NewNumber(8),
			data.NewNumber(9),
			data.NewNumber(10),
		),
	},
}
