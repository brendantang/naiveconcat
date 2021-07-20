package parse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	wordRunes            = "abcdefghijjklmnopqrstuvwxyz+-*/" // wordRunes are legal characters in a word name.
	stringDelimiter rune = '"'                               // delimiter that opens or closes a string.
	stringEscaper   rune = '\\'                              // delimiter that escapes the following character in a string.
	quotationOpener rune = '{'                               // delimiter that opens a quotation.
	quotationCloser rune = '}'                               // delimiter that closes a quotation.
	eof             rune = -1
)

// the lexer scans an input src and emits a stream of lexed tokens.
type lexer struct {
	src      string     // source text to lex.
	out      chan token // where lexed tokens are sent.
	errs     chan error // where lexing errors are sent.
	done     chan bool  // send true when done lexing src.
	behavior lexingFn   // function defining lexing behavior.
	startPos int        // selection start position.
	endPos   int        // selection end position.
}

func newLexer(src string, start lexingFn) *lexer {
	return &lexer{
		src:      src,
		out:      make(chan token, 2),
		errs:     make(chan error, 1),
		done:     make(chan bool, 1),
		behavior: start,
	}
}

type lexingFn func(l *lexer) lexingFn

func (l *lexer) run() {
	for l.behavior != nil {
		l.behavior = l.behavior(l)
	}
}

func lexMain(l *lexer) lexingFn {
	for r, width := l.peek(); r != eof; r, width = l.peek() {
		switch {
		case r == '-' || r == '.': // either could be a word or the beginning of a number.
			if nextR, _ := l.peek(); !unicode.IsDigit(nextR) {
				return lexWord
			}
			fallthrough

		case isNumeric(r):
			return lexNumber

		case isWhitespace(r):
			l.ignore(width)

		case matchRune(wordRunes, r):
			return lexWord

		case r == stringDelimiter:
			l.ignore(width) // ignore the opening "
			return lexString

		case r == quotationOpener:
			l.endPos += width
			l.commit(openQ)

		case r == quotationCloser:
			l.endPos += width
			l.commit(closeQ)

		default:
			return nil
		}
	}
	return nil
}

func lexNumber(l *lexer) lexingFn {
	l.acceptOne("-.")
	l.accept("0123456789")
	l.acceptOne(".")
	l.accept("0123456789")
	l.commit(num)
	return lexMain
}

func lexWord(l *lexer) lexingFn {
	l.accept(wordRunes)
	l.commit(word)
	return lexMain
}

func lexString(l *lexer) lexingFn {
	var (
		current  rune
		previous rune
		width    int
	)
	// an unescaped " terminates the string
	shouldTerminate := func() bool {
		return (current == stringDelimiter && previous != stringEscaper)
	}
	for !shouldTerminate() {
		previous = current
		current, width = l.peek()
		if !shouldTerminate() {
			l.endPos += width
		}
	}
	l.commit(str)
	l.ignore(width)
	return lexMain
}

// peek returns the next rune without adding it to the selection.
func (l *lexer) peek() (r rune, width int) {
	if l.endPos >= len(l.src) {
		return eof, 0
	}
	return utf8.DecodeRuneInString(l.src[l.endPos:])
}

// next returns the next rune and includes it in the selection.
func (l *lexer) next() (r rune) {
	// return EOF if reached end of source
	if l.endPos >= len(l.src) {
		return eof
	}

	r, width := utf8.DecodeRuneInString(l.src[l.endPos:])
	l.endPos += width
	return r
}

// commit uses the selection and given type to initialize a token and sends it
// to the out channel. The beginning of the next selection starts at the end of
// the previous.
func (l *lexer) commit(typ tokenType) {
	l.out <- token{typ, l.selection()}
	l.startPos = l.endPos
}

// selection returns the slice of the input string between the startPos
// and endPos.
func (l *lexer) selection() string {
	return l.src[l.startPos:l.endPos]
}

func (l *lexer) accept(valid string) {
	for l.acceptOne(valid) {
	}
}
func (l *lexer) acceptOne(valid string) bool {
	r, width := l.peek()
	if !(strings.IndexRune(valid, r) >= 0) {
		return false
	}
	l.endPos += width
	return true
}

func (l *lexer) ignore(width int) {
	l.endPos += width
	l.startPos = l.endPos
}

func isWhitespace(r rune) bool {
	return matchRune(" \n\r\t", r)
}

func isNumeric(r rune) bool {
	return matchRune("0123456789-+.", r)
}

func matchRune(valid string, r rune) bool {
	return strings.IndexRune(valid, r) >= 0
}
