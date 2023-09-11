package ast

import (
	"gci/token"
	"testing"
)

func Test_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar1337"},
					Value: "myVar1337",
				},
			},
		},
	}

	if program.String() != "let myVar = myVar1337;" {
		t.Errorf("program.String() wronmg. got=%q", program.String())
	}
}
