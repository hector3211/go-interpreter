package ast

import (
	"go-interpreter/internal/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVal"},
					Value: "anotherVal",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVal;" {
		t.Errorf("program.String() failed, got: %q", program.String())
	}
}
