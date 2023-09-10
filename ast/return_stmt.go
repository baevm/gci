package ast

import "gci/token"

type ReturnStatement struct {
	Token       token.Token // 'return' token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
