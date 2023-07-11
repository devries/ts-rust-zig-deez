package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{Let, "let"},
		{Ident, "five"},
		{Assign, "="},
		{Int, "5"},
		{Semicolon, ";"},
		{Let, "let"},
		{Ident, "ten"},
		{Assign, "="},
		{Int, "10"},
		{Semicolon, ";"},
		{Let, "let"},
		{Ident, "add"},
		{Assign, "="},
		{Function, "fn"},
		{LParen, "("},
		{Ident, "x"},
		{Comma, ","},
		{Ident, "y"},
		{RParen, ")"},
		{LSquirly, "{"},
		{Ident, "x"},
		{Plus, "+"},
		{Ident, "y"},
		{Semicolon, ";"},
		{RSquirly, "}"},
		{Semicolon, ";"},
		{Let, "let"},
		{Ident, "result"},
		{Assign, "="},
		{Ident, "add"},
		{LParen, "("},
		{Ident, "five"},
		{Comma, ","},
		{Ident, "ten"},
		{RParen, ")"},
		{Semicolon, ";"},
		{Bang, "!"},
		{Minus, "-"},
		{ForwardSlash, "/"},
		{Asterisk, "*"},
		{Int, "5"},
		{Semicolon, ";"},
		{Int, "5"},
		{LessThan, "<"},
		{Int, "10"},
		{GreaterThan, ">"},
		{Int, "5"},
		{Semicolon, ";"},
		{If, "if"},
		{LParen, "("},
		{Int, "5"},
		{LessThan, "<"},
		{Int, "10"},
		{RParen, ")"},
		{LSquirly, "{"},
		{Return, "return"},
		{True, "true"},
		{Semicolon, ";"},
		{RSquirly, "}"},
		{Else, "else"},
		{LSquirly, "{"},
		{Return, "return"},
		{False, "false"},
		{Semicolon, ";"},
		{RSquirly, "}"},
		{Int, "10"},
		{Equal, "=="},
		{Int, "10"},
		{Semicolon, ";"},
		{Int, "10"},
		{NotEqual, "!="},
		{Int, "9"},
		{Semicolon, ";"},
		{Eof, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
