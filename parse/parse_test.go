package parse

import (
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestTokenize(t *testing.T) {
	for _, c := range testCases {
		got := tokenize(c.have)
		if len(got) != len(c.wantTokens) {
			t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n", c.description, c.wantTokens, got)
		}
		for i, tok := range got {
			if tok != c.wantTokens[i] {
				t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n", c.description, c.wantTokens, got)
			}
		}
	}
}

func TestParse(t *testing.T) {
	for _, c := range testCases {
		got, err := Parse(c.have)
		if err != nil {
			t.Fatalf("FAIL: %s\nParsing err: %v", c.description, err)
		}
		if len(got) != len(c.wantTokens) {
			t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n", c.description, c.wantValues, got)
		}
		for i, val := range got {
			if val.String() != c.wantValues[i].String() {
				t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n", c.description, c.wantValues, got)
			}
		}
	}
}

var testCases = []struct {
	description string
	have        string
	wantTokens  []token
	wantValues  []data.Value
}{
	{
		description: "numbers",
		have:        "1 2 3 54.3",
		wantTokens: []token{
			{num, "1"},
			{num, "2"},
			{num, "3"},
			{num, "54.3"},
		},
		wantValues: []data.Value{
			data.NewNumber(1),
			data.NewNumber(2),
			data.NewNumber(3),
			data.NewNumber(54.3),
		},
	},
	{
		description: "numbers and words",
		have:        "1 2 3 foo 3",
		wantTokens: []token{
			{num, "1"},
			{num, "2"},
			{num, "3"},
			{word, "foo"},
			{num, "3"},
		},
		wantValues: []data.Value{
			data.NewNumber(1),
			data.NewNumber(2),
			data.NewNumber(3),
			data.NewWord("foo"),
			data.NewNumber(3),
		},
	},
	{
		description: "string",
		have:        `"foo" "2" "string with spaces"`,
		wantTokens: []token{
			{str, "foo"},
			{str, "2"},
			{str, "string with spaces"},
		},
		wantValues: []data.Value{
			data.NewString("foo"),
			data.NewString("2"),
			data.NewString("string with spaces"),
		},
	},
	{
		description: "quotation",
		have:        "1 { 2 3 }",
		wantTokens: []token{
			{num, "1"},
			{openQ, "{"},
			{num, "2"},
			{num, "3"},
			{closeQ, "}"},
		},
		wantValues: []data.Value{
			data.NewNumber(1),
			data.NewQuotation(
				data.NewNumber(2),
				data.NewNumber(3),
			),
		},
	},
}
