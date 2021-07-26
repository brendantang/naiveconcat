package parse

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"strconv"
)

type Parser struct {
	in   chan token      // where lexed tokens are received.
	Out  chan data.Value // where parsed expressions are sent.
	Errs chan error      // where parsing errors are sent.
}

func NewParser(input chan token) *Parser {
	return &Parser{
		in:   input,
		Out:  make(chan data.Value, 2),
		Errs: make(chan error, 1),
	}

}

func (p *Parser) Run() {
loop:
	for tok := range p.in {
		switch tok.typ {
		case num:
			if n, err := strconv.ParseFloat(tok.body, 64); err != nil {
				p.Errs <- conversionError(tok, data.Number)
			} else {
				p.Out <- data.NewNumber(n)
			}
		case word:
			p.Out <- data.NewWord(tok.body)
		case str:
			p.Out <- data.NewString(tok.body)
		case openQ:
			subParser := NewParser(p.in)
			go subParser.Run()
			var quotedVals []data.Value
			for more := true; more; {
				select {
				case subVal, ok := <-subParser.Out:
					if !ok {
						more = ok
						break
					}
					quotedVals = append(quotedVals, subVal)

				case err := <-subParser.Errs:
					if err != nil {
						p.Errs <- err
					}
				}
			}
			p.Out <- data.NewQuotation(quotedVals...)
		case closeQ:
			break loop
		}
	}
	close(p.Out)
	close(p.Errs)
}

func conversionError(tok token, typ data.Type) error {
	return fmt.Errorf("could not parse %v as %v", tok, typ)
}
