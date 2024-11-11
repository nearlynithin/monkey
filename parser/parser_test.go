package parser

import (
	"log"
	"testing"

	"github.com/nearlynithin/monkey/ast"
	"github.com/nearlynithin/monkey/lexer"
)

//Test cases for the let statementp
func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 282883;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		log.Fatalf("ParseProgram() function returned nil")
	}
	if len(program.Statements) != 3 {
		log.Fatalf("program.Statements doe not contain 3 statements. got=%d",len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},{"y"},{"foobar"},
	}
	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt,tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral()!= "let" {
		t.Errorf("s.TokenLiteral() not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}
	
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got=%s",name,letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got=%s",name,letStmt.Name)
		return false
	}

	return true
}