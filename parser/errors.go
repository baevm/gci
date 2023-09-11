package parser

import (
	"fmt"
	"gci/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(tkn token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tkn, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(tkn token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", tkn)
	p.errors = append(p.errors, msg)
}
