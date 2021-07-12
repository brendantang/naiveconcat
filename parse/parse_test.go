package parse

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	for _, c := range parseTestCases {
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

var parseTestCases = []struct {
	description string
	have        string
	wantTokens  []token
}{
	{
		description: "numbers",
		have:        "1 2 3 54.3",
		wantTokens:  []token{"1", "2", "3", "54.3"},
	},
}
