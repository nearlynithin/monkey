package lexer

import "github.com/nearlynithin/monkey/token"



type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}


func New(input string) *Lexer { //method to return a lexer
	l := &Lexer{input: input}
	l.readChar()
	return l
}


func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition;
	l.readPosition+=1
}

//return a Token struct
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l* Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=' :
		tok = newToken(token.ASSIGN, l.ch)
	case ';' :
		tok = newToken(token.SEMICOLON, l.ch)
	case '(' :
		tok = newToken(token.LPAREN, l.ch)
	case ')' :
		tok = newToken(token.RPAREN, l.ch)
	case ',' :
		tok = newToken(token.COMMA, l.ch)
	case '+' :
		tok = newToken(token.PLUS, l.ch)
	case '{' :
		tok = newToken(token.LBRACE, l.ch)
	case '}' :
		tok = newToken(token.RBRACE, l.ch)
	case 0: //This represents ASCII for null value
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}