package parser

import (
	"fmt"
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
)

type decoder interface {
	decode() ast.Expression
}

type parseExpressionFn func(precedence int) ast.Expression
type parseBlockStmtFn func() *ast.BlockStmt

type infixParseFn func(ast.Expression) ast.Expression

type tokenDecoderMap map[token.TokenType]decoder
type infixParserMap map[token.TokenType]infixParseFn

type Parser struct {
	scanner
	decoderMap    tokenDecoderMap
	infixParseFns infixParserMap
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{scanner: scanner{l: l, errors: []string{}}}

	identifierDecoder := &identifier{&p.scanner}
	integerDecoder := &integer{&p.scanner}
	booleanDecoder := &boolean{&p.scanner}
	prefixExprDecoder := &prefixExpr{&p.scanner, p.parseExpression}
	groupedExprDecoder := &groupedExpr{&p.scanner, p.parseExpression}
	ifExprDecoder := &ifExpr{&p.scanner, p.parseExpression, p.parseBlockStmt}

	p.decoderMap = tokenDecoderMap{
		token.IDENT:  identifierDecoder,
		token.INT:    integerDecoder,
		token.TRUE:   booleanDecoder,
		token.FALSE:  booleanDecoder,
		token.NOT:    prefixExprDecoder,
		token.SUB:    prefixExprDecoder,
		token.LPAREN: groupedExprDecoder,
		token.IF:     ifExprDecoder,
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

func (this *Parser) parseExpression(precedence int) ast.Expression {
	tokenDecoder := this.decoderMap[this.curTok.Type]
	if nil == tokenDecoder {
		this.appendError(fmt.Sprintf("%v has no decoder", token.ToString(this.curTok.Type)))
		return nil
	}
	leftExpr := tokenDecoder.decode()

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
