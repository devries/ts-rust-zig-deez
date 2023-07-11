package monkey

import (
	"fmt"
	"unicode"
)

// TokenType is the type of token.
type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	IDENT
	INT
	ASSIGN
	PLUS
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	default:
		return "Unknown"
	}
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Token is a struct for a token.
type Token struct {
	Type    TokenType
	Literal string
}

// String is the token stringer
func (t Token) String() string {
	return fmt.Sprintf("{Type:%s Literal:%q}", t.Type, t.Literal)
}

// Lexer is a lexer struct.
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

// NewLexer creates a new lexer from a string input.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{ASSIGN, "="}
	case ';':
		tok = Token{SEMICOLON, ";"}
	case '(':
		tok = Token{LPAREN, "("}
	case ')':
		tok = Token{RPAREN, ")"}
	case ',':
		tok = Token{COMMA, ","}
	case '+':
		tok = Token{PLUS, "+"}
	case '{':
		tok = Token{LBRACE, "{"}
	case '}':
		tok = Token{RBRACE, "}"}
	case 0:
		tok = Token{EOF, ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = Token{LookupIdent(literal), literal}
			return tok
		} else if isDigit(l.ch) {
			number := l.readNumber()
			tok = Token{INT, number}
			return tok
		} else {
			tok = Token{ILLEGAL, string(l.ch)}
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '_'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// LookupIdent returns the token type from string.
func LookupIdent(ident string) TokenType {
	if tt, ok := keywords[ident]; ok {
		return tt
	}

	return IDENT
}
