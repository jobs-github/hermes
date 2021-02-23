package parser

import (
	"fmt"
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  *token.Token
	peekTok *token.Token
	errors  []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// init curTok & peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (this *Parser) Errors() []string {
	return this.errors
}

func (this *Parser) peekError(t token.TokenType) {
	this.errors = append(this.errors, fmt.Sprintf("expected next token to be %v, got %v instead", token.ToString(t), token.ToString(this.peekTok.Type)))
}

func (this *Parser) nextToken() {
	this.curTok = this.peekTok
	this.peekTok = this.l.NextToken()
}

func (this *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: ast.StatementSlice{}}
	for !this.curTok.Eof() {
		stmt := this.parseStmt()
		if nil != stmt {
			program.Statements = append(program.Statements, stmt)
		}
		this.nextToken()
	}
	return program
}

func (this *Parser) parseStmt() ast.Statement {
	switch this.curTok.Type {
	case token.VAR:
		return this.parseVarStmt()
	default:
		return nil
	}
}

func (this *Parser) expectPeek(t token.TokenType) bool {
	if this.peekTok.TypeIs(t) {
		this.nextToken()
		return true
	}
	this.peekError(t)
	return false
}

func (this *Parser) parseVarStmt() ast.Statement {
	stmt := &ast.VarStatement{Tok: this.curTok}
	if !this.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Tok: this.curTok, Value: this.curTok.Literal}
	if !this.expectPeek(token.ASSIGN) {
		return nil
	}

	for !this.curTok.TypeIs(token.SEMICOLON) {
		this.nextToken()
	}
	return stmt
}
