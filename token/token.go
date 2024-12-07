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
	SLASH = "/"
	BANG = "!"
	MINUS = "-"
	ASTERISK = "*"
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	//delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	STRING = "STRING"
)

var keywords = map[string]TokenType {
	"fn" : FUNCTION,
	"let" : LET,
	"if" : IF,
	"else" : ELSE,
	"true" : TRUE,
	"false" : FALSE,
	"return" : RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok,ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}