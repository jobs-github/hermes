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

func (this *Lexer) twoCharToken(tokenType token.TokenType, expectedNextChar byte, tokenType2 token.TokenType, literal string) *token.Token {
	if expectedNextChar == this.peekChar() {
		this.readChar()
		return &token.Token{Type: tokenType2, Literal: literal}
	} else {
		return newToken(tokenType, this.ch)
	}
}

func (this *Lexer) NextToken() *token.Token {
	var tok *token.Token
	this.skipWhitespace()

	if 0 == this.ch {
		return &token.Token{Type: token.EOF, Literal: ""}
	}

	switch this.ch {
	case '+':
		tok = this.twoCharToken(token.ADD, '+', token.INC, "++")
	case '-':
		tok = this.twoCharToken(token.SUB, '-', token.DEC, "--")
	case '&':
		tok = this.twoCharToken(token.ILLEGAL, '&', token.AND, "&&")
	case '|':
		tok = this.twoCharToken(token.ILLEGAL, '|', token.OR, "||")
	case '=':
		tok = this.twoCharToken(token.ASSIGN, '=', token.EQ, "==")
	case '!':
		tok = this.twoCharToken(token.NOT, '=', token.NEQ, "!=")
	case '<':
		tok = this.twoCharToken(token.LT, '=', token.LEQ, "<=")
	case '>':
		tok = this.twoCharToken(token.GT, '=', token.GEQ, ">=")
	default:
		tt, ok := token.GetTokenType(this.ch)
		if ok {
			tok = newToken(tt, this.ch)
		} else {
			if isLetter(this.ch) {
				literal := this.readIdentifier()
				return &token.Token{Type: token.LookupIdent(literal), Literal: literal}
			} else if isDigit(this.ch) {
				return &token.Token{Type: token.INT, Literal: this.readNumber()}
			} else {
				tok = newToken(token.ILLEGAL, this.ch)
			}
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

func (this *Lexer) peekChar() byte {
	if this.nextPosition >= len(this.input) {
		return 0
	} else {
		return this.input[this.nextPosition]
	}
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
