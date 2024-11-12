package ast

import (
	"testing"

	"github.com/nearlynithin/monkey/token"
)

func testString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name : &Identifier{Token : token.Token{Type: token.IDENT, Literal: "myVar"}, Value : "myVar",},
				Value : &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar",},			
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("proprogram.String() is wrong, got=%q",program.String())
	}
}
