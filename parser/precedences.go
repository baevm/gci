package parser

import (
	"gci/token"
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS, // Lowest precedence
	token.NOT_EQ:   EQUALS,
	token.LT:       LTGT,
	token.GT:       LTGT,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL, // Highest precedence
}

func (p *Parser) peekPrecedece() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}

	return LOWEST
}
