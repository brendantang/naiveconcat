package parse

import (
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
			{tNum, "1"},
			{tNum, "2"},
			{tNum, "3"},
			{tNum, "54.3"},
		},
	},
	{
		description: "numbers and words",
		have:        "1 2 3 foo 3",
		wantTokens: []token{
			{tNum, "1"},
			{tNum, "2"},
			{tNum, "3"},
			{tWord, "foo"},
			{tNum, "3"},
		},
	},
	{
		description: "string",
		have:        `"foo" "2" "string with spaces"`,
		wantTokens: []token{
			{tStr, "foo"},
			{tStr, "2"},
			{tStr, "string with spaces"},
		},
	},
	{
		description: "quotation",
		have:        "1 { 2 3 }",
		wantTokens: []token{
			{tNum, "1"},
			{tOpenQ, "{"},
			{tNum, "2"},
			{tNum, "3"},
			{tCloseQ, "}"},
		},
	},
}
