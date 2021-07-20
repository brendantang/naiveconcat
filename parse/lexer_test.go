package parse

import (
	"testing"
)

func TestLexer(t *testing.T) {
	for _, c := range testCases {

		l := newLexer(c.src, lexMain)
		go l.run()

		var got []token
		var err error

		for done := false; done; {
			select {
			case tok := <-l.out:
				got = append(got, tok)

			case err = <-l.errs:
				if err != nil {
					t.Fatalf(
						"FAIL: %s\nLexing error: %v",
						c.description,
						err,
					)
				}
			case done = <-l.done:
				break
			}
		}

		if len(got) != len(c.wantTokens) {
			failLexTest(t, c, got, "length doesn't match")
		}

		for i, tok := range got {
			if tok != c.wantTokens[i] {
				failLexTest(t, c, got, "mismatched elements")
			}
		}
	}
}

func failLexTest(t *testing.T, c testCase, got []token, msg string) {
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.wantTokens, got, msg)
}
