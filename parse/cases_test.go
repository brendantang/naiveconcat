package parse

import (
	"github.com/brendantang/naiveconcat/data"
)

type testCase struct {
	description string
	src         string
	wantTokens  []token
	wantValues  []data.Value
}

var testCases = []testCase{
	{"a number", "23", []token{{num, "23"}}, []data.Value{data.NewNumber(23)}},
	{"negative number", "-1000000", []token{{num, "-1000000"}}, []data.Value{data.NewNumber(-1000000)}},
	{"decimal number", "1000.333", []token{{num, "1000.333"}}, []data.Value{data.NewNumber(1000.333)}},
	{
		"multiple numbers",
		"23 -11 23.003 -12.32",
		[]token{
			{num, "23"},
			{num, "-11"},
			{num, "23.003"},
			{num, "-12.32"},
		},
		[]data.Value{
			data.NewNumber(23),
			data.NewNumber(-11),
			data.NewNumber(23.003),
			data.NewNumber(-12.32),
		},
	},
	{
		"numbers and words",
		"1 2 3 foo 3",
		[]token{
			{num, "1"},
			{num, "2"},
			{num, "3"},
			{word, "foo"},
			{num, "3"},
		},
		[]data.Value{
			data.NewNumber(1),
			data.NewNumber(2),
			data.NewNumber(3),
			data.NewWord("foo"),
			data.NewNumber(3),
		},
	},
	{
		"string",
		`"foo" "2" "string with spaces"`,
		[]token{
			{str, "foo"},
			{str, "2"},
			{str, "string with spaces"},
		},
		[]data.Value{
			data.NewString("foo"),
			data.NewString("2"),
			data.NewString("string with spaces"),
		},
	},
	{
		"operators",
		"- + / *",
		[]token{
			{word, "-"},
			{word, "+"},
			{word, "/"},
			{word, "*"},
		},
		[]data.Value{
			data.NewWord("-"),
			data.NewWord("+"),
			data.NewWord("/"),
			data.NewWord("*"),
		},
	},
	{
		"whitespace",
		"foo \n bar \r \t baz",
		[]token{
			{word, "foo"},
			{word, "bar"},
			{word, "baz"},
		},
		[]data.Value{
			data.NewWord("foo"),
			data.NewWord("bar"),
			data.NewWord("baz"),
		},
	},
	{
		"quotation",
		"1 { 2 3 }",
		[]token{
			{num, "1"},
			{openQ, "{"},
			{num, "2"},
			{num, "3"},
			{closeQ, "}"},
		},
		[]data.Value{
			data.NewNumber(1),
			data.NewQuotation(
				data.NewNumber(2),
				data.NewNumber(3),
			),
		},
	},
	{
		"comments",
		`1 2 -- comment begins
		--full line comment
		"foo"`,
		[]token{
			{num, "1"},
			{num, "2"},
			{str, "foo"},
		},
		[]data.Value{
			data.NewNumber(1),
			data.NewNumber(2),
			data.NewString("foo"),
		},
	},
	{
		"nested quotation",
		"1 { 2 { 3 } }",
		[]token{
			{num, "1"},
			{openQ, "{"},
			{num, "2"},
			{openQ, "{"},
			{num, "3"},
			{closeQ, "}"},
			{closeQ, "}"},
		},
		[]data.Value{
			data.NewNumber(1),
			data.NewQuotation(
				data.NewNumber(2),
				data.NewQuotation(
					data.NewNumber(3),
				),
			),
		},
	},
}
