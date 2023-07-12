package ast

import (
	"monkey/lexer"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: lexer.Token{Type: lexer.Let, Literal: "let"},
				Name: &Identifier{
					Token: lexer.Token{Type: lexer.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: lexer.Token{Type: lexer.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
