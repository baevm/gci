package ast

import "gci/token"

type LetStatement struct {
	Token token.Token // 'let' token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
