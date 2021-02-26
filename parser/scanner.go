package parser

import (
	"fmt"
	"hermes/lexer"
	"hermes/token"
)

type scanner struct {
	l       *lexer.Lexer
	curTok  *token.Token
	peekTok *token.Token
	errors  []string
}

func (this *scanner) peekPrecedence() int {
	return getPrecedence(this.peekTok)
}

func (this *scanner) curPrecedence() int {
	return getPrecedence(this.curTok)
}

func (this *scanner) Errors() []string {
	return this.errors
}

func (this *scanner) appendError(err string) {
	this.errors = append(this.errors, err)
}

func (this *scanner) nextToken() {
	this.curTok = this.peekTok
	this.peekTok = this.l.NextToken()
}

func (this *scanner) expectPeek(t token.TokenType) bool {
	if this.peekTok.TypeIs(t) {
		this.nextToken()
		return true
	}
	this.appendError(fmt.Sprintf("expected next token to be %v, got %v instead", token.ToString(t), token.ToString(this.peekTok.Type)))
	return false
}
