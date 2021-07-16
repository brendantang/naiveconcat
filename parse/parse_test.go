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

var testCases = []struct {
	description string
	have        string
	wantTokens  []token
}{
	{
		description: "numbers",
		have:        "1 2 3 54.3",
		wantTokens: []token{
			{data.Number, "1"},
			{data.Number, "2"},
			{data.Number, "3"},
			{data.Number, "54.3"},
		},
	},
	{
		description: "numbers and words",
		have:        "1 2 3 foo 3",
		wantTokens: []token{
			{data.Number, "1"},
			{data.Number, "2"},
			{data.Number, "3"},
			{data.Word, "foo"},
			{data.Number, "3"},
		},
	},
	{
		description: "string",
		have:        `"foo" "2" "string with spaces"`,
		wantTokens: []token{
			{data.String, "foo"},
			{data.String, "2"},
			{data.String, "string with spaces"},
		},
	},
	{
		description: "quotation",
		have:        "1 { 2 3 }",
		wantTokens: []token{
			{data.Number, "1"},
			{data.Quotation, "{ 2 3 }"},
		},
	},
}
