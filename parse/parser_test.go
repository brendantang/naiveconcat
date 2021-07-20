package parse

import (
	"github.com/brendantang/naiveconcat/data"
	"testing"
)

func TestParser(t *testing.T) {
	for _, c := range testCases {
		var in = make(chan token, 1)
		p := NewParser(in)

		go func() {
			for _, tok := range c.wantTokens {
				in <- tok
			}
			close(in)
		}()
		go p.Run()

		var got []data.Value

		for more := true; more; {
			select {
			case val, ok := <-p.Out:
				if !ok {
					more = ok
					break
				}

				got = append(got, val)
			case err := <-p.Errs:
				if err != nil {
					t.Fatalf("FAIL: %s\nParsing error: %v", c.description, err)
				}
			}
		}

		if len(got) != len(c.wantValues) {
			failParseTest(t, c, got, "length doesn't match")
		}

		for i, val := range got {
			if val.String() != c.wantValues[i].String() {
				failParseTest(t, c, got, "mismatched elements")
			}
		}

	}
}

func failParseTest(t *testing.T, c testCase, got []data.Value, msg string) {
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.wantValues, got, msg)
}
