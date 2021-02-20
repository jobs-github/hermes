package parser

import (
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  *token.Token
	peekTok *token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// init curTok & peekTok
	p.nextTok()
	p.nextTok()
	return p
}

func (this *Parser) nextTok() {
	this.curTok = this.peekTok
	this.peekTok = this.l.NextToken()
}

func (this *Parser) ParseProgram() *ast.Program {
	return nil
}
