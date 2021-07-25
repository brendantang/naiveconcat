package parse

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	wordRunes            = "abcdefghijjklmnopqrstuvwxyz+-*/=" // wordRunes are legal characters in a word name.
	stringDelimiter rune = '"'                                // delimiter that opens or closes a string.
	stringEscaper   rune = '\\'                               // delimiter that escapes the following character in a string.
	quotationOpener rune = '{'                                // delimiter that opens a quotation.
	quotationCloser rune = '}'                                // delimiter that closes a quotation.
)

// Lexer scans an input src and emits a stream of lexed tokens.
type Lexer struct {
	src      string     // source text to lex.
	Out      chan token // where lexed tokens are sent.
	Errs     chan error // where lexing errors are sent.
	Done     chan bool
	behavior lexingFn    // function defining lexing behavior.
	startPos int         // selection start position.
	endPos   int         // selection end position.
	Debug    *log.Logger // where to print out lexer Debugging info.
}

// NewLexer returns a *Lexer with initialized Out, Errs, and Done channels.
func NewLexer(src string) *Lexer {
	return &Lexer{
		src:      src,
		Out:      make(chan token, 1),
		Done:     make(chan bool, 1),
		Errs:     make(chan error, 1),
		behavior: lexMain,
		Debug:    nil,
	}
}

type lexingFn func(l *Lexer) lexingFn

// Run scans the source text and emits tokens on the Out channel. When an error
// is encountered, it is sent on the Errs channel. When lexing is finished,
// true is sent on the Done channel.
func (l *Lexer) Run() {
	for l.behavior != nil {
		l.behavior = l.behavior(l)
	}
	close(l.Out)
	close(l.Errs)
	l.Done <- true
	close(l.Done)
}

func lexMain(l *Lexer) lexingFn {
	for r, width := l.peek(); r != EOF; r, width = l.peek() {
		switch {
		case r == '-' || r == '.': // either could be a word or the beginning of a number.
			switch nextR, _ := l.runeAt(l.endPos + width); {
			case nextR == '-': // '--' indicates a comment
				return lexComment
			case !unicode.IsDigit(nextR):
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

func lexComment(l *Lexer) lexingFn {
	if l.Debug != nil {
		l.Debug.Println("lexing comment")
	}
	l.skipUntil("\n\r")
	return lexMain
}

func lexNumber(l *Lexer) lexingFn {
	l.acceptOne("-.")
	l.accept("0123456789")
	l.acceptOne(".")
	l.accept("0123456789")
	l.commit(num)
	return lexMain
}

func lexWord(l *Lexer) lexingFn {
	l.accept(wordRunes)
	l.commit(word)
	return lexMain
}

func lexString(l *Lexer) lexingFn {
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
func (l *Lexer) peek() (r rune, width int) {
	if l.endPos >= len(l.src) {
		r = EOF
		width = 0
	} else {
		r, width = l.runeAt(l.endPos)
	}
	if l.Debug != nil {
		l.Debug.Printf("peek: %s\n", string(r))
	}
	return
}

// next returns the next rune and includes it in the selection.
func (l *Lexer) next() (r rune) {
	// return EOF if reached end of source
	if l.endPos >= len(l.src) {
		return EOF
	}

	r, width := l.peek()
	l.endPos += width
	return r
}

func (l *Lexer) runeAt(pos int) (r rune, width int) {
	return utf8.DecodeRuneInString(l.src[pos:])
}

// commit uses the selection and given type to initialize a token and sends it
// to the out channel. The beginning of the next selection starts at the end of
// the previous.
func (l *Lexer) commit(typ tokenType) {
	tok := token{typ, l.selection()}
	l.Out <- tok

	l.startPos = l.endPos
}

// selection returns the slice of the input string between the startPos
// and endPos.
func (l *Lexer) selection() string {
	if l.startPos >= l.endPos {
		l.Errs <- l.errorf("tried to take a selection that starts at %d but ends at %d", l.startPos, l.endPos)
		return ""
	}
	return l.src[l.startPos:l.endPos]
}

func (l *Lexer) accept(valid string) {
	for l.acceptOne(valid) {
	}
}
func (l *Lexer) acceptOne(valid string) bool {
	r, width := l.peek()
	if !(strings.IndexRune(valid, r) >= 0) {
		return false
	}
	l.endPos += width
	return true
}

func (l *Lexer) ignore(width int) {
	l.endPos += width
	l.startPos = l.endPos
	if l.Debug != nil {
		l.Debug.Printf("lexing Debug: ignore start:%d, end:%d\n", l.startPos, l.endPos)
	}
}

func (l *Lexer) skipUntil(want string) {
	for r, width := l.peek(); !(strings.IndexRune(want, r) >= 0); r, width = l.peek() {
		l.ignore(width)
	}
}

func isWhitespace(r rune) bool {
	return matchRune(" \n\r\t", r)
}

func isNumeric(r rune) bool {
	return matchRune("0123456789-.", r)
}

func matchRune(valid string, r rune) bool {
	return strings.IndexRune(valid, r) >= 0
}

func (l *Lexer) errorf(format string, args ...interface{}) error {
	format = "lexing error at position %d: '%s': " + format
	args = append([]interface{}{l.startPos, l.src[l.startPos:l.endPos]}, args...)
	return fmt.Errorf(format, args...)
}
