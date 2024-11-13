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
	checkParserErrors(t,p)
	if program == nil {
		log.Fatalf("ParseProgram() function returned nil")
	}
	if len(program.Statements) != 3 {
		log.Fatalf("program.Statements does not contain 3 statements. got=%d",len(program.Statements))
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

//method to log errors if any
func checkParserErrors(t *testing.T, p* Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors",len(errors))
	for _,msg := range errors {
		t.Errorf("parse error: %q",msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 7283479;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",len(program.Statements))
	}

	for _,stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T",stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got=%q",returnStmt.TokenLiteral())
		}
	}
}

//Parser tests
func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements! got=%d",len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T",stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not %s, got=%s","foobar",ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s. got=%s","foobar",ident.TokenLiteral())
	}

}

//test for integer expressions
func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l:= lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments, got=%d",len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an *ast.ExpressionStatement, got=%T",program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.IntegerLiteral, got=%T",stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value is not %d, got=%d",5,literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral is not %s, got=%d","5",literal.TokenLiteral())
	}
}