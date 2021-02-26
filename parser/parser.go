package parser

import (
	"fmt"
	"hermes/ast"
	"hermes/lexer"
	"hermes/token"
)

type decodeInfix func(ast.Expression) ast.Expression

type infixDecoderMap map[token.TokenType]decodeInfix

type Parser struct {
	scanner

	stmtParser    *stmtParser
	tokenDecoders tokenDecoderMap
	infixDecoders infixDecoderMap
}

func newInfixDecoders(parseInfixExpr decodeInfix) infixDecoderMap {
	return infixDecoderMap{
		token.LT:  parseInfixExpr,
		token.GT:  parseInfixExpr,
		token.ADD: parseInfixExpr,
		token.SUB: parseInfixExpr,
		token.MUL: parseInfixExpr,
		token.DIV: parseInfixExpr,
		token.MOD: parseInfixExpr,
		token.EQ:  parseInfixExpr,
		token.NEQ: parseInfixExpr,
		token.LEQ: parseInfixExpr,
		token.GEQ: parseInfixExpr,
		token.AND: parseInfixExpr,
		token.OR:  parseInfixExpr,
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{scanner: scanner{l: l, errors: []string{}}}

	p.stmtParser = newStmtParser(&p.scanner, p.parseExpression)
	p.tokenDecoders = newTokenDecoders(&p.scanner, p.parseExpression, p.parseBlockStmt)
	p.infixDecoders = newInfixDecoders(p.parseInfixExpression)
	// init curTok & peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (this *Parser) parseStmt() ast.Statement {
	return this.stmtParser.decode(this.curTok.Type)
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

func (this *Parser) parseExpression(precedence int) ast.Expression {
	tokenDecoder := this.tokenDecoders[this.curTok.Type]
	if nil == tokenDecoder {
		this.appendError(fmt.Sprintf("%v has no decoder", token.ToString(this.curTok.Type)))
		return nil
	}
	leftExpr := tokenDecoder.decode()

	for !this.peekTok.TypeIs(token.SEMICOLON) && precedence < this.peekPrecedence() {
		infix := this.infixDecoders[this.peekTok.Type]
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
