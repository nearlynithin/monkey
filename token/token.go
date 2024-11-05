package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

//the token types
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identifiers + literals
	IDENT = "IDENT" // variable names like add,foo,bar etc
	INT = "INT" //integers like 1,2324,0

	//operators
	ASSIGN = "="
	PLUS = "+"

	//delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	
)