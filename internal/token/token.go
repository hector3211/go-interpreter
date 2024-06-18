package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDNET" // add,foobar,x,y,..
	INT   TokenType = "INT"   // 12345

	// Operators
	ASSIGN      TokenType = "="
	EQUAL       TokenType = "=="
	NOTEQUAL    TokenType = "!="
	BANG        TokenType = "!"
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	SLASH       TokenType = "/"
	ASTERIK     TokenType = "*"
	LESSTHAN    TokenType = "<"
	GREATERTHAN TokenType = ">"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LSQUIRLY TokenType = "{"
	RSQUIRLY TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	TRUE     TokenType = "TRUE"
	FLASE    TokenType = "FALSE"
	RETURN   TokenType = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
	// FileName string
	// Reader  io.Reader
	Err error
}

func CreateToken(token TokenType, char byte) Token {
	return Token{
		Type:    token,
		Literal: string(char),
	}
}

var keyWords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FLASE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyWords[ident]; ok {
		return tok
	}
	return IDENT
}
