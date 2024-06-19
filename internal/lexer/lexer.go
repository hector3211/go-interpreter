package lexer

import "go-interpreter/internal/token"

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = token.CreateToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOTEQUAL, Literal: literal}
		} else {
			tok = token.CreateToken(token.BANG, l.ch)
		}
	case ';':
		tok = token.CreateToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.CreateToken(token.LPAREN, l.ch)
	case ')':
		tok = token.CreateToken(token.RPAREN, l.ch)
	case ',':
		tok = token.CreateToken(token.COMMA, l.ch)
	case '+':
		tok = token.CreateToken(token.PLUS, l.ch)
	case '-':
		tok = token.CreateToken(token.MINUS, l.ch)
	case '*':
		tok = token.CreateToken(token.ASTERIK, l.ch)
	case '/':
		tok = token.CreateToken(token.SLASH, l.ch)
	case '{':
		tok = token.CreateToken(token.LSQUIRLY, l.ch)
	case '}':
		tok = token.CreateToken(token.RSQUIRLY, l.ch)
	case '<':
		tok = token.CreateToken(token.LESSTHAN, l.ch)
	case '>':
		tok = token.CreateToken(token.GREATERTHAN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.CreateToken(token.ILLEGAL, l.ch)
		}
	}
	l.ReadChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
