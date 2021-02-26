package lexer

import (
	"testing"

	"hermes/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `var five = 5;
	var ten = 10;
	var add = func(x, y) {
		x + y;
	};
	var result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	10 == 10;
	10 != 9;
	true && false;
	true || false;
	`

	tests := []struct {
		name        string
		wantType    token.TokenType
		wantLiteral string
	}{
		{"1", token.VAR, "var"},
		{"2", token.IDENT, "five"},
		{"3", token.ASSIGN, "="},
		{"4", token.INT, "5"},
		{"5", token.SEMICOLON, ";"},
		{"6", token.VAR, "var"},
		{"7", token.IDENT, "ten"},
		{"8", token.ASSIGN, "="},
		{"9", token.INT, "10"},
		{"10", token.SEMICOLON, ";"},
		{"11", token.VAR, "var"},
		{"12", token.IDENT, "add"},
		{"13", token.ASSIGN, "="},
		{"14", token.FUNC, "func"},
		{"15", token.LPAREN, "("},
		{"16", token.IDENT, "x"},
		{"17", token.COMMA, ","},
		{"18", token.IDENT, "y"},
		{"19", token.RPAREN, ")"},
		{"20", token.LBRACE, "{"},
		{"21", token.IDENT, "x"},
		{"22", token.ADD, "+"},
		{"23", token.IDENT, "y"},
		{"24", token.SEMICOLON, ";"},
		{"25", token.RBRACE, "}"},
		{"26", token.SEMICOLON, ";"},
		{"27", token.VAR, "var"},
		{"28", token.IDENT, "result"},
		{"29", token.ASSIGN, "="},
		{"30", token.IDENT, "add"},
		{"31", token.LPAREN, "("},
		{"32", token.IDENT, "five"},
		{"33", token.COMMA, ","},
		{"34", token.IDENT, "ten"},
		{"35", token.RPAREN, ")"},
		{"36", token.SEMICOLON, ";"},
		{"37", token.NOT, "!"},
		{"38", token.SUB, "-"},
		{"39", token.DIV, "/"},
		{"40", token.MUL, "*"},
		{"41", token.INT, "5"},
		{"42", token.SEMICOLON, ";"},
		{"43", token.INT, "5"},
		{"44", token.LT, "<"},
		{"45", token.INT, "10"},
		{"46", token.GT, ">"},
		{"47", token.INT, "5"},
		{"48", token.SEMICOLON, ";"},
		{"49", token.IF, "if"},
		{"50", token.LPAREN, "("},
		{"51", token.INT, "5"},
		{"52", token.LT, "<"},
		{"53", token.INT, "10"},
		{"54", token.RPAREN, ")"},
		{"55", token.LBRACE, "{"},
		{"56", token.RETURN, "return"},
		{"57", token.TRUE, "true"},
		{"58", token.SEMICOLON, ";"},
		{"59", token.RBRACE, "}"},
		{"60", token.ELSE, "else"},
		{"61", token.LBRACE, "{"},
		{"62", token.RETURN, "return"},
		{"63", token.FALSE, "false"},
		{"64", token.SEMICOLON, ";"},
		{"65", token.RBRACE, "}"},
		{"66", token.INT, "10"},
		{"67", token.EQ, "=="},
		{"68", token.INT, "10"},
		{"69", token.SEMICOLON, ";"},
		{"70", token.INT, "10"},
		{"71", token.NEQ, "!="},
		{"72", token.INT, "9"},
		{"73", token.SEMICOLON, ";"},
		{"74", token.TRUE, "true"},
		{"75", token.AND, "&&"},
		{"76", token.FALSE, "false"},
		{"77", token.SEMICOLON, ";"},
		{"78", token.TRUE, "true"},
		{"79", token.OR, "||"},
		{"80", token.FALSE, "false"},
		{"81", token.SEMICOLON, ";"},
		{"EOF", token.EOF, ""},
	}

	l := New(input)
	for _, tt := range tests {
		tok := l.nextToken()
		if tok.Type != tt.wantType {
			t.Fatalf("[%v], Lexer.NextToken() type = %v, want %v", tt.name, tok.Type, tt.wantType)
		}
		if tok.Literal != tt.wantLiteral {
			t.Fatalf("[%v], Lexer.NextToken() literal = %v, want %v", tt.name, tok.Literal, tt.wantLiteral)
		}
	}
}
