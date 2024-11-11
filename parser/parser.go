package parser

import (
	"github.com/nearlynithin/monkey/ast"
	"github.com/nearlynithin/monkey/lexer"
	"github.com/nearlynithin/monkey/token"
)



type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read token twice, first time peekToken is set, next curToken is set 
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} //creating a root node
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET :
		return p.parseLetStatment()
	default:
		return nil
	}
}

//called when current token is LET
func (p* Parser) parseLetStatment() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken} //making a node off of token.LET

	if !p.expectPeek(token.IDENT) {
		return nil //if IDENT was not found
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil //if ASSIGN was not found
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken() // loop until SEMICOLON is reached
	}

	return stmt
}

func (p *Parser) curTokenIs( t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs( t token.TokenType) bool {
	return p.peekToken.Type == t
}


func (p *Parser) expectPeek( t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken() // if got the expected token, update current token immediately
		return true
	} else {
		return false
	}
}