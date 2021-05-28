package parser

import (
	"Q/ast"
	"Q/token"
)

func stmtEnd(scanner *scanner) bool {
	return scanner.curTok.TypeIs(token.SEMICOLON) || scanner.curTok.Eof()
}

type stmtDecoder interface {
	decode() ast.Statement
}

type stmtParser struct {
	scanner       *scanner
	assignDecoder stmtDecoder
	exprDecoder   stmtDecoder
	m             map[token.TokenType]stmtDecoder
}

func (this *stmtParser) isAssignStmt() bool {
	return this.scanner.expectCurPeek(token.IDENT, token.ASSIGN)
}

func (this *stmtParser) decode(t token.TokenType) ast.Statement {
	decoder, ok := this.m[t]
	if ok {
		return decoder.decode()
	}
	if this.isAssignStmt() {
		return this.assignDecoder.decode()
	}
	return this.exprDecoder.decode()
}

func newStmtParser(s *scanner, parseExpression parseExpressionFn) *stmtParser {
	return &stmtParser{
		scanner:       s,
		assignDecoder: &assignStmt{s, parseExpression},
		exprDecoder:   &exprStmt{s, parseExpression},
		m: map[token.TokenType]stmtDecoder{
			token.VAR:    &varStmt{s, parseExpression},
			token.RETURN: &returnStmt{s, parseExpression},
			token.BREAK:  &breakStmt{s},
		},
	}
}

// varStmt : implement stmtDecoder
type varStmt struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *varStmt) decode() ast.Statement {
	stmt := &ast.VarStmt{Tok: this.scanner.curTok}
	if !this.scanner.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Tok: this.scanner.curTok, Value: this.scanner.curTok.Literal}
	if !this.scanner.expectPeek(token.ASSIGN) {
		return nil
	}

	this.scanner.nextToken()

	stmt.Value = this.parseExpression(PRECED_LOWEST)

	for !stmtEnd(this.scanner) {
		this.scanner.nextToken()
	}
	return stmt
}

// returnStmt : implement stmtDecoder
type returnStmt struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *returnStmt) decode() ast.Statement {
	stmt := &ast.ReturnStmt{Tok: this.scanner.curTok}
	this.scanner.nextToken()

	stmt.ReturnValue = this.parseExpression(PRECED_LOWEST)

	for !stmtEnd(this.scanner) {
		this.scanner.nextToken()
	}
	return stmt
}

// exprStmt : implement stmtDecoder
type exprStmt struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *exprStmt) decode() ast.Statement {
	stmt := &ast.ExpressionStmt{Tok: this.scanner.curTok}
	stmt.Expr = this.parseExpression(PRECED_LOWEST)

	if this.scanner.peekTok.TypeIs(token.SEMICOLON) {
		this.scanner.nextToken()
	}
	return stmt
}

// breakStmt : implement stmtDecoder
type breakStmt struct {
	scanner *scanner
}

func (this *breakStmt) decode() ast.Statement {
	return &ast.BreakStmt{Tok: this.scanner.curTok}
}

// assignStmt : implement stmtDecoder
type assignStmt struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *assignStmt) decode() ast.Statement {
	stmt := &ast.AssignStmt{}
	stmt.Name = &ast.Identifier{Tok: this.scanner.curTok, Value: this.scanner.curTok.Literal}
	if !this.scanner.expectPeek(token.ASSIGN) {
		return nil
	}

	this.scanner.nextToken()

	stmt.Value = this.parseExpression(PRECED_LOWEST)

	for !stmtEnd(this.scanner) {
		this.scanner.nextToken()
	}
	return stmt
}
