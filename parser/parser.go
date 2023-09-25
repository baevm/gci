package parser

import (
	"gci/ast"
	"gci/lexer"
	"gci/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token // points to current token
	peekToken token.Token // points to next token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn // called when encounter token in prefix
	infixParseFns  map[token.TokenType]infixParseFn  // called when encounter token in infix
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Higher const = bigger precedence over last operator
const (
	LOWEST  = iota + 1
	EQUALS  // ==
	LTGT    // > <
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // function(x)
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read 2 tokens to set currToken and peekToken
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) registerPrefix(tkn token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tkn] = fn
}

func (p *Parser) registerInfix(tkn token.TokenType, fn infixParseFn) {
	p.infixParseFns[tkn] = fn
}

func (p *Parser) expectPeek(tkn token.TokenType) bool {
	if p.peekTokenIs(tkn) {
		p.nextToken()
		return true
	} else {
		p.peekError(tkn)
		return false
	}
}

func (p *Parser) currTokenIs(tkn token.TokenType) bool {
	return p.currToken.Type == tkn
}

func (p *Parser) peekTokenIs(tkn token.TokenType) bool {
	return p.peekToken.Type == tkn
}
