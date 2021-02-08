package lexer

import (
	"github.com/jobs-github/hermes/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch byte) *token.Token {
	return &token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '_' == ch
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isWhitespace(c byte) bool {
	return c == ' ' ||
		c == '\t' ||
		c == '\n' ||
		c == '\v' ||
		c == '\f' ||
		c == '\r'
}

func (this *Lexer) skipWhitespace() {
	for isWhitespace(this.ch) {
		this.readChar()
	}
}

func (this *Lexer) NextToken() *token.Token {
	var tok *token.Token
	this.skipWhitespace()
	switch this.ch {
	case '=':
		tok = newToken(token.ASSIGN, this.ch)
	case ';':
		tok = newToken(token.SEMICOLON, this.ch)
	case '(':
		tok = newToken(token.LPAREN, this.ch)
	case ')':
		tok = newToken(token.RPAREN, this.ch)
	case ',':
		tok = newToken(token.COMMA, this.ch)
	case '+':
		tok = newToken(token.ADD, this.ch)
	case '{':
		tok = newToken(token.LBRACE, this.ch)
	case '}':
		tok = newToken(token.RBRACE, this.ch)
	case 0:
		return &token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(this.ch) {
			literal := this.readIdentifier()
			return &token.Token{Type: token.LookupIdent(literal), Literal: literal}
		} else if isDigit(this.ch) {
			return &token.Token{Type: token.INT, Literal: this.readNumber()}
		} else {
			tok = newToken(token.ILLEGAL, this.ch)
		}
	}
	this.readChar()
	return tok
}

func (this *Lexer) readNumber() string {
	pos := this.position
	for isDigit(this.ch) {
		this.readChar()
	}
	return this.input[pos:this.position]
}

func (this *Lexer) readIdentifier() string {
	pos := this.position
	for isLetter(this.ch) {
		this.readChar()
	}
	return this.input[pos:this.position]
}

func (this *Lexer) readChar() {
	if this.nextPosition >= len(this.input) {
		this.ch = 0
	} else {
		this.ch = this.input[this.nextPosition]
	}
	this.position = this.nextPosition
	this.nextPosition = this.nextPosition + 1
}
