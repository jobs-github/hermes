package parser

import (
	"fmt"
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
	"strconv"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

type prefixParserMap map[token.TokenType]prefixParseFn
type infixParserMap map[token.TokenType]infixParseFn

type Parser struct {
	scanner
	prefixParseFns prefixParserMap
	infixParseFns  infixParserMap
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{scanner: scanner{l: l, errors: []string{}}}
	p.prefixParseFns = prefixParserMap{
		token.IDENT:  p.parseIdentifier,
		token.INT:    p.parseIntegerLiteral,
		token.TRUE:   p.parseBoolean,
		token.FALSE:  p.parseBoolean,
		token.NOT:    p.parsePrefixExpression,
		token.SUB:    p.parsePrefixExpression,
		token.LPAREN: p.parseGroupedExpression,
		token.IF:     p.parseIfExpression,
	}
	p.infixParseFns = infixParserMap{
		token.LT:  p.parseInfixExpression,
		token.GT:  p.parseInfixExpression,
		token.ADD: p.parseInfixExpression,
		token.SUB: p.parseInfixExpression,
		token.MUL: p.parseInfixExpression,
		token.DIV: p.parseInfixExpression,
		token.MOD: p.parseInfixExpression,
		token.EQ:  p.parseInfixExpression,
		token.NEQ: p.parseInfixExpression,
		token.LEQ: p.parseInfixExpression,
		token.GEQ: p.parseInfixExpression,
		token.AND: p.parseInfixExpression,
		token.OR:  p.parseInfixExpression,
	}
	// init curTok & peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (this *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Stmts: ast.StatementSlice{}}
	for !this.curTok.Eof() {
		stmt := this.parseStmt()
		if nil != stmt {
			program.Stmts = append(program.Stmts, stmt)
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

func (this *Parser) parseBlockStmt() *ast.BlockStmt {
	block := &ast.BlockStmt{Tok: this.curTok}
	block.Stmts = ast.StatementSlice{}
	this.nextToken()
	for !this.curTok.TypeIs(token.RBRACE) {
		stmt := this.parseStmt()
		if nil != stmt {
			block.Stmts = append(block.Stmts, stmt)
		}
		this.nextToken()
	}
	return block
}

func (this *Parser) parseVarStmt() ast.Statement {
	stmt := &ast.VarStmt{Tok: this.curTok}
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
	stmt := &ast.ReturnStmt{Tok: this.curTok}
	this.nextToken()

	for !this.curTok.TypeIs(token.SEMICOLON) {
		this.nextToken()
	}
	return stmt
}

func (this *Parser) parseExprStmt() ast.Statement {
	stmt := &ast.ExpressionStmt{Tok: this.curTok}
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
	expr := &ast.Integer{Tok: this.curTok}
	val, err := strconv.ParseInt(this.curTok.Literal, 0, 64)
	if nil != err {
		this.appendError(fmt.Sprintf("could not parse %v as integer", this.curTok.Literal))
		return nil
	}
	expr.Value = val
	return expr
}

func (this *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Tok: this.curTok, Value: this.curTok.TypeIs(token.TRUE)}
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
	expr.Right = this.parseExpression(PRECED_PREFIX)
	return expr
}

func (this *Parser) parseGroupedExpression() ast.Expression {
	this.nextToken()
	expr := this.parseExpression(PRECED_LOWEST)
	if !this.expectPeek(token.RPAREN) {
		return nil
	}
	return expr
}

func (this *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Tok: this.curTok, Clauses: ast.IfClauseSlice{}}

	clause := &ast.IfClause{}
	if !this.expectPeek(token.LPAREN) {
		return nil
	}
	this.nextToken()
	clause.If = this.parseExpression(PRECED_LOWEST)
	if !this.expectPeek(token.RPAREN) {
		return nil
	}
	if !this.expectPeek(token.LBRACE) {
		return nil
	}
	clause.Then = this.parseBlockStmt()
	expr.Clauses = append(expr.Clauses, clause)

	if this.peekTok.TypeIs(token.ELSE) {
		this.nextToken()
		if !this.expectPeek(token.LBRACE) {
			return nil
		}
		expr.Else = this.parseBlockStmt()
	}

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
