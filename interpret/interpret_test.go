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
			failInterpretTest(t, c, tmp, "error opening temporary file", err.Error())
		}
		defer os.Remove(tmp.Name()) // defer cleanup
		os.Stdout = tmp

		// interpret the test case
		d, s := eval.CoreDict(), data.NewStack()
		for i, line := range c.src {
			err = Interpret(line+"\r", d, s)
			t.Log(line)
			if err != nil {
				failInterpretTest(t, c, tmp, fmt.Sprintf("interpreter error on line %d", i+1), err.Error())
			}
		}

		output, err := ioutil.ReadFile(tmp.Name())
		if err != nil {
			failInterpretTest(t, c, tmp, "error reading captured output", err.Error())
		}

		got := strings.Trim(string(output), "\n\r \t")
		if got != c.want {
			failInterpretTest(t, c, tmp, got, "")
		}

		// set stdout back to normal
		os.Stdout = normalStdout
	}
}

func failInterpretTest(t *testing.T, c interpretTestCase, tmp *os.File, got string, msg string) {
	capturedOutput, _ := ioutil.ReadAll(tmp)
	t.Log(capturedOutput)
	t.Fatalf("FAIL: %s\nWant: %s\nGot: %s\n%s\n\n", c.description, c.want, got, msg)
}

type interpretTestCase struct {
	description string
	src         []string
	want        string
}
