package lexer

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
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	EQ
	NOT_EQ

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
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
	case MINUS:
		return "-"
	case BANG:
		return "!"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case LT:
		return "<"
	case GT:
		return ">"
	case EQ:
		return "=="
	case NOT_EQ:
		return "!="
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
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	default:
		return "Unknown"
	}
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
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
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{EQ, "=="}
		} else {
			tok = Token{ASSIGN, "="}
		}
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
	case '-':
		tok = Token{MINUS, "-"}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{NOT_EQ, "!="}
		} else {
			tok = Token{BANG, "!"}
		}
	case '*':
		tok = Token{ASTERISK, "*"}
	case '/':
		tok = Token{SLASH, "/"}
	case '<':
		tok = Token{LT, "<"}
	case '>':
		tok = Token{GT, ">"}
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

func (l *Lexer) peekChar() rune {
	return l.input[l.readPosition]
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
