package lexer

import (
	"fmt"
	"unicode"
)

//go:generate stringer -type=TokenType

// TokenType is the type of token.
type TokenType int

const (
	Illegal TokenType = iota
	Eof

	Ident
	Int

	Assign
	Plus
	Minus
	Bang
	Asterisk
	ForwardSlash

	LessThan
	GreaterThan
	Equal
	NotEqual

	Comma
	Semicolon

	LParen
	RParen
	LSquirly
	RSquirly

	Function
	Let
	True
	False
	If
	Else
	Return
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
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
			tok = Token{Equal, "=="}
		} else {
			tok = Token{Assign, "="}
		}
	case ';':
		tok = Token{Semicolon, ";"}
	case '(':
		tok = Token{LParen, "("}
	case ')':
		tok = Token{RParen, ")"}
	case ',':
		tok = Token{Comma, ","}
	case '+':
		tok = Token{Plus, "+"}
	case '-':
		tok = Token{Minus, "-"}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{NotEqual, "!="}
		} else {
			tok = Token{Bang, "!"}
		}
	case '*':
		tok = Token{Asterisk, "*"}
	case '/':
		tok = Token{ForwardSlash, "/"}
	case '<':
		tok = Token{LessThan, "<"}
	case '>':
		tok = Token{GreaterThan, ">"}
	case '{':
		tok = Token{LSquirly, "{"}
	case '}':
		tok = Token{RSquirly, "}"}
	case 0:
		tok = Token{Eof, ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = Token{LookupIdent(literal), literal}
			return tok
		} else if isDigit(l.ch) {
			number := l.readNumber()
			tok = Token{Int, number}
			return tok
		} else {
			tok = Token{Illegal, string(l.ch)}
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

	return Ident
}
