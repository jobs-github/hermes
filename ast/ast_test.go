package ast

import (
	"hermes/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Stmts: StatementSlice{
			&VarStmt{
				Tok: &token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Tok:   &token.Token{Type: token.IDENT, Literal: "testVar1"},
					Value: "testVar1",
				},
				Value: &Identifier{
					Tok:   &token.Token{Type: token.IDENT, Literal: "testVar2"},
					Value: "testVar2",
				},
			},
		},
	}
	if program.String() != "var testVar1 = testVar2;" {
		t.Errorf("program.String() wrong, got: %v", program.String())
	}
}
