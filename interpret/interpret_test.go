package interpret

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"github.com/brendantang/naiveconcat/eval"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func testInterpret(t *testing.T, cases []interpretTestCase) {
	for _, c := range cases {

		// capture stdout during the test case
		normalStdout := os.Stdout
		tmp, err := ioutil.TempFile("", "captured_output")
		if err != nil {
			failInterpretTest(t, c, "error opening temporary file", err.Error())
		}
		defer os.Remove(tmp.Name()) // defer cleanup
		os.Stdout = tmp

		// interpret the test case
		d, s := eval.CoreDict(), data.NewStack()
		for i, line := range c.src {
			t.Log(s)
			err = Interpret(line+"\r", d, s)
			if err != nil {
				failInterpretTest(t, c, fmt.Sprintf("interpreter error on line %d: %s", i+1, line), err.Error())
			}
		}

		output, err := ioutil.ReadFile(tmp.Name())
		if err != nil {
			failInterpretTest(t, c, "error reading captured output", err.Error())
		}

		got := strings.Trim(string(output), "\n\r \t")
		if got != c.want {
			failInterpretTest(t, c, got, "")
		}

		// set stdout back to normal
		os.Stdout = normalStdout
	}
}

func failInterpretTest(t *testing.T, c interpretTestCase, got string, msg string) {
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.want, got, msg)
}

type interpretTestCase struct {
	description string
	src         []string
	want        string
}
