package lexer

import (
	"miguenz-lang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let one = 1;
	let two = 2;

	let add = fn(x, y) {
		x + y
	};

	let result = add(one, two);
	!-/*1;
	1 < 2 > 1

	if (1 < 2) {
		return true;
	} else {
		return false;
	}

	4 == 4;
	5 != 3;

	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "one"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "one"},
		{token.COMMA, ","},
		{token.IDENT, "two"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},

		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "2"},
		{token.GT, ">"},
		{token.INT, "1"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "4"},
		{token.EQ, "=="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.NOT_EQ, "!="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong for char: [%q]. expected=%q, got=%q",
				i, tok.Literal, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong for char: [%q]. expected=%q, got=%q",
				i, tok.Literal, tt.expectedLiteral, tok.Literal)
		}
	}
}
