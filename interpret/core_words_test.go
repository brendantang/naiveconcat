package interpret

import (
	"testing"
)

func TestInterpretCoreWords(t *testing.T) {
	testInterpret(t, coreWordTestCases)
}

func TestWords(t *testing.T) {
}

var coreWordTestCases = []interpretTestCase{
	{
		"say",
		`"foo" "bar" say say
		1.234567 say
		{5 +} say
		{5 +} lambda say
		`,
		"\"bar\"\n\"foo\"\n1.234567\n{5 +}\nPROCEDURE",
	},
	{
		"stack",
		`"foo" "bar" 1.234567 {5 +} {5 +} lambda
		stack
		`,
		"[\"foo\" \"bar\" 1.234567 {5 +} PROCEDURE]",
	},
	{
		"words",
		`
		3 "three" let
		words
		`,
		`*	PROCEDURE
+	PROCEDURE
-	PROCEDURE
/	PROCEDURE
=	PROCEDURE
and	PROCEDURE
append	PROCEDURE
apply	PROCEDURE
define	PROCEDURE
drop	PROCEDURE
dup	PROCEDURE
false	FALSE
join	PROCEDURE
lambda	PROCEDURE
length	PROCEDURE
let	PROCEDURE
lop	PROCEDURE
not	PROCEDURE
or	PROCEDURE
say	PROCEDURE
split	PROCEDURE
stack	PROCEDURE
then	PROCEDURE
three	3
true	TRUE
words	PROCEDURE`,
	},
	{
		"words with a local dict",
		`{"foo" "x" let words} apply`,
		`x	"foo"`,
	},
	{
		"append to a quotation",
		`4 {1 2 3} append say`,
		`{1 2 3 4}`,
	},
}
