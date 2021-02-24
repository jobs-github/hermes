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
	PRECED_LOWEST
	PRECED_OR     // ||
	PRECED_AND    // &&
	PRECED_EQ     // ==
	PRECED_NEQ    // != < > >= <=
	PRECED_ADD    // +
	PRECED_MUL    // *
	PRECED_PREFIX // -x !x
	PRECED_CALL   // myFn(x)
)

var (
	precedences = map[token.TokenType]int{
		token.LT: PRECED_NEQ,
		token.GT: PRECED_NEQ,
		// ASSIGN
		// NOT
		token.ADD: PRECED_ADD,
		token.SUB: PRECED_ADD,
		token.MUL: PRECED_MUL,
		token.DIV: PRECED_MUL,
		token.MOD: PRECED_MUL,
		token.EQ:  PRECED_EQ,
		token.NEQ: PRECED_NEQ,
		token.LEQ: PRECED_NEQ,
		token.GEQ: PRECED_NEQ,
		token.AND: PRECED_AND,
		token.OR:  PRECED_OR,
	}
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

type prefixParserMap map[token.TokenType]prefixParseFn
type infixParserMap map[token.TokenType]infixParseFn

type Parser struct {
	l       *lexer.Lexer
	curTok  *token.Token
	peekTok *token.Token
	errors  []string

	prefixParseFns prefixParserMap
	infixParseFns  infixParserMap
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFns = prefixParserMap{
		token.IDENT: p.parseIdentifier,
		token.INT:   p.parseIntegerLiteral,
		token.NOT:   p.parsePrefixExpression,
		token.SUB:   p.parsePrefixExpression,
	}
	p.infixParseFns = infixParserMap{
		token.ADD: p.parseInfixExpression,
		token.SUB: p.parseInfixExpression,
		token.MUL: p.parseInfixExpression,
		token.DIV: p.parseInfixExpression,
		token.MOD: p.parseInfixExpression,
		token.EQ:  p.parseInfixExpression,
		token.NEQ: p.parseInfixExpression,
		token.LT:  p.parseInfixExpression,
		token.GT:  p.parseInfixExpression,
	}
	// init curTok & peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func getPrecedence(tok *token.Token) int {
	if v, ok := precedences[tok.Type]; ok {
		return v
	}
	return PRECED_LOWEST
}

func (this *Parser) peekPrecedence() int {
	return getPrecedence(this.peekTok)
}

func (this *Parser) curPrecedence() int {
	return getPrecedence(this.curTok)
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
	stmt.Expr = this.parseExpression(PRECED_LOWEST)

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
		this.appendError(fmt.Sprintf("%v has no prefix fn", token.ToString(this.curTok.Type)))
		return nil
	}
	leftExpr := prefixFn()

	for !this.peekTok.TypeIs(token.SEMICOLON) && precedence < this.peekPrecedence() {
		infix := this.infixParseFns[this.peekTok.Type]
		if nil == infix {
			return leftExpr
		}
		this.nextToken()
		leftExpr = infix(leftExpr)
	}
	return leftExpr
}

func (this *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Tok: this.curTok,
		Op:  this.curTok.Literal,
	}
	this.nextToken()
	expr.Right = this.parseExpression(PRECED_LOWEST)
	return expr
}

func (this *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Tok:  this.curTok,
		Op:   this.curTok.Literal,
		Left: left,
	}
	preced := this.curPrecedence()
	this.nextToken()
	expr.Right = this.parseExpression(preced)
	return expr
}
