package parse

import (
	"github.com/brendantang/naiveconcat/data"
	"strings"
	"unicode/utf8"
)

const (
	wordRunes            = "abcdefghijjklmnopqrstuvwxyz-" // wordRunes are legal characters in a word name.
	stringDelimiter rune = '"'                            // delimiter that opens or closes a string.
	stringEscaper   rune = '\\'                           // delimiter that escapes the following character in a string.
)

type lexer struct {
	source         string
	selectionStart int
	selectionEnd   int
	tokens         []token
	behavior       lexingFn
}

type lexingFn func(l *lexer) lexingFn

func (l *lexer) run() {
	for l.behavior != nil {
		l.behavior = l.behavior(l)
	}
}

// default behavior determines a more specific lexingFn to define behavior.
func defaultBehavior(l *lexer) lexingFn {
	for r, width := l.peek(); r != eof; r, width = l.peek() {
		switch {
		case isNumeric(r):
			return lexNumber
		case isWhitespace(r):
			l.ignore(width)
		case matchRune(wordRunes, r):
			return lexWord
		case r == stringDelimiter:
			l.ignore(width) // ignore the opening "
			return lexString
		default:
			return nil
		}
	}
	return nil
}

func lexNumber(l *lexer) lexingFn {
	l.acceptOne("+-.")
	l.accept("0123456789")
	l.acceptOne(".")
	l.accept("0123456789")
	l.commit(data.Number)
	return defaultBehavior
}

func lexWord(l *lexer) lexingFn {
	l.accept(wordRunes)
	l.commit(data.Word)
	return defaultBehavior
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
			l.selectionEnd += width
		}
	}
	l.commit(data.String)
	l.ignore(width)
	return defaultBehavior
}

// peek returns the next rune without adding it to the selection.
func (l *lexer) peek() (r rune, width int) {
	if l.selectionEnd >= len(l.source) {
		return eof, 0
	}
	return utf8.DecodeRuneInString(l.source[l.selectionEnd:])
}

// next returns the next rune and includes it in the selection.
func (l *lexer) next() (r rune) {
	// return EOF if reached end of source
	if l.selectionEnd >= len(l.source) {
		return eof
	}

	r, width := utf8.DecodeRuneInString(l.source[l.selectionEnd:])
	l.selectionEnd += width
	return r
}

// commit uses the selection and given type to add a new token to tokens. The
// beginning of the next selection starts at the end of the previous.
func (l *lexer) commit(t data.Type) {
	tok := token{t, l.selection()}
	l.tokens = append(l.tokens, tok)
	l.selectionStart = l.selectionEnd
}

// selection returns the slice of the input string between the selectionStart
// and selectionEnd.
func (l *lexer) selection() string {
	return l.source[l.selectionStart:l.selectionEnd]
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
	l.selectionEnd += width
	return true
}

func (l *lexer) ignore(width int) {
	l.selectionEnd += width
	l.selectionStart = l.selectionEnd
}

const eof rune = -1

func isWhitespace(r rune) bool {
	return matchRune(" \n\r\t", r)
}

func isNumeric(r rune) bool {
	return matchRune("0123456789-+.", r)
}

func matchRune(valid string, r rune) bool {
	return strings.IndexRune(valid, r) >= 0
}
