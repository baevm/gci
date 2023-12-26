package lexer

import "gci/token"

type Lexer struct {
	input        string
	position     int  // current position (current char)
	readPosition int  // current reading position (after current char)
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// check if we are at the end of input
	// if so set char to EOF
	// else move to next position
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tkn token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tkn = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '-':
		tkn = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tkn = newToken(token.BANG, l.ch)
		}
	case '*':
		tkn = newToken(token.ASTERISK, l.ch)
	case '/':
		tkn = newToken(token.SLASH, l.ch)
	case '<':
		tkn = newToken(token.LT, l.ch)
	case '>':
		tkn = newToken(token.GT, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case '{':
		tkn = newToken(token.LBRACE, l.ch)
	case '}':
		tkn = newToken(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	case '"':
		tkn.Type = token.STRING
		tkn.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		} else if isDigit(l.ch) {
			tkn.Literal = l.readNumber()
			tkn.Type = token.INT
			return tkn
		} else {
			tkn = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tkn
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// function to check allowed letters in identifier
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// function to check if char is digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// function to skip whitespaces in input
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// function that gets identifier name from input
// let |var| = 5
func (l *Lexer) readIdentifier() string {
	currentPosition := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[currentPosition:l.position]
}

func (l *Lexer) readNumber() string {
	currentPosition := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[currentPosition:l.position]
}

// function to read whole string
func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}
