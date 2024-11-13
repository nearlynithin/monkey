package parser

import (
	"fmt"

	"github.com/nearlynithin/monkey/ast"
	"github.com/nearlynithin/monkey/lexer"
	"github.com/nearlynithin/monkey/token"
)

const (
	_ int = iota
	LOWEST 
	EQUALS 		//==
	LESSGREATER //< or >
	SUM 		//+
	PRODUCT 	//*
	PREFIX		//-X or !X
	CALL		// myFunction(X)
)


type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
	errors[] string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn

}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

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
	case token.RETURN :
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatment()
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

		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",t,p.peekToken.Type)
	p.errors = append(p.errors, msg)
}


//called in case of a RETURN token
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken() // skip all expressions untill SEMICOLON is reached
	}

	return stmt
}

type ( 
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

//Helper functions to add entries to the maps
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatment() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}