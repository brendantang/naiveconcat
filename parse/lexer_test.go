package parse

import (
	"testing"
)

func TestLexer(t *testing.T) {
	for _, c := range testCases {

		l := NewLexer(c.src)
		l.Run()

		var got []token
		var err error

		for more := true; more; {
			t.Logf("%#v", l)
			select {
			case tok, ok := <-l.Out:
				t.Log("received from out", tok)
				if false {
				}
				got = append(got, tok)
				more = ok

			case err = <-l.Errs:
				if err != nil {
					t.Fatalf(
						"FAIL: %s\nLexing error: %v",
						c.description,
						err,
					)
				}
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
