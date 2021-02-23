package parser

import (
	"fmt"
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQ     // ==
	LTGT   // < >
	ADD    // +
	MUL    // *
	PREFIX // -x !x
	CALL   // myFn(x)
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

type Parser struct {
	l       *lexer.Lexer
	curTok  *token.Token
	peekTok *token.Token
	errors  []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFns = map[token.TokenType]prefixParseFn{
		token.IDENT: p.parseIdentifier,
		token.INT:   p.parseIntegerLiteral,
	}
	// init curTok & peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (this *Parser) Errors() []string {
	return this.errors
}

func (this *Parser) appendError(err string) {
	this.errors = append(this.errors, err)
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
	case token.RETURN:
		return this.parseReturnStmt()
	default:
		return this.parseExprStmt()
	}
}

func (this *Parser) expectPeek(t token.TokenType) bool {
	if this.peekTok.TypeIs(t) {
		this.nextToken()
		return true
	}
	this.appendError(fmt.Sprintf("expected next token to be %v, got %v instead", token.ToString(t), token.ToString(this.peekTok.Type)))
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

func (this *Parser) parseReturnStmt() ast.Statement {
	stmt := &ast.ReturnStatement{Tok: this.curTok}
	this.nextToken()

	for !this.curTok.TypeIs(token.SEMICOLON) {
		this.nextToken()
	}
	return stmt
}

func (this *Parser) parseExprStmt() ast.Statement {
	stmt := &ast.ExpressionStatement{Tok: this.curTok}
	stmt.Expr = this.parseExpression(LOWEST)

	if this.peekTok.TypeIs(token.SEMICOLON) {
		this.nextToken()
	}
	return stmt
}

func (this *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Tok: this.curTok, Value: this.curTok.Literal}
}

func (this *Parser) parseIntegerLiteral() ast.Expression {
	expr := &ast.IntegerLiteral{Tok: this.curTok}
	val, err := strconv.ParseInt(this.curTok.Literal, 0, 64)
	if nil != err {
		this.appendError(fmt.Sprintf("could not parse %v as integer", this.curTok.Literal))
		return nil
	}
	expr.Value = val
	return expr
}

func (this *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn := this.prefixParseFns[this.curTok.Type]
	if nil == prefixFn {
		return nil
	}
	leftExpr := prefixFn()
	return leftExpr
}
