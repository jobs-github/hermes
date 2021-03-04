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
	scanner       *scanner
	stmtParser    *stmtParser
	tokenDecoders tokenDecoderMap
	infixDecoders infixDecoderMap
}

func newInfixDecoders(parseInfixExpr decodeInfix, parseCall decodeInfix) infixDecoderMap {
	return infixDecoderMap{
		token.LT:     parseInfixExpr,
		token.GT:     parseInfixExpr,
		token.ADD:    parseInfixExpr,
		token.SUB:    parseInfixExpr,
		token.MUL:    parseInfixExpr,
		token.DIV:    parseInfixExpr,
		token.MOD:    parseInfixExpr,
		token.EQ:     parseInfixExpr,
		token.NEQ:    parseInfixExpr,
		token.LEQ:    parseInfixExpr,
		token.GEQ:    parseInfixExpr,
		token.AND:    parseInfixExpr,
		token.OR:     parseInfixExpr,
		token.LPAREN: parseCall,
	}
}

func New(l *lexer.Lexer) (*Parser, error) {
	s, err := newScanner(l)
	if nil == s {
		return nil, err
	}
	p := &Parser{scanner: s}
	p.stmtParser = newStmtParser(s, p.parseExpression)
	p.tokenDecoders = newTokenDecoders(s, p.parseExpression, p.parseBlockStmt)
	p.infixDecoders = newInfixDecoders(p.parseInfixExpression, p.parseCallExpression)
	return p, nil
}

func (this *Parser) Errors() []string {
	return this.scanner.errors
}

func (this *Parser) parseStmt() ast.Statement {
	return this.stmtParser.decode(this.scanner.curTok.Type)
}

func (this *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Stmts: ast.StatementSlice{}}
	for !this.scanner.eof() {
		stmt := this.parseStmt()
		if nil != stmt {
			program.Stmts = append(program.Stmts, stmt)
		}
		this.scanner.nextToken()
	}
	return program
}

func (this *Parser) parseBlockStmt() *ast.BlockStmt {
	block := &ast.BlockStmt{Tok: this.scanner.curTok}
	block.Stmts = ast.StatementSlice{}
	this.scanner.nextToken()
	for !this.scanner.curTok.TypeIs(token.RBRACE) {
		stmt := this.parseStmt()
		if nil != stmt {
			block.Stmts = append(block.Stmts, stmt)
		}
		this.scanner.nextToken()
	}
	return block
}

func (this *Parser) parseExpression(precedence int) ast.Expression {
	tokenDecoder := this.tokenDecoders[this.scanner.curTok.Type]
	if nil == tokenDecoder {
		this.scanner.appendError(fmt.Sprintf("%v has no decoder", token.ToString(this.scanner.curTok.Type)))
		return nil
	}
	leftExpr := tokenDecoder.decode()

	for !this.scanner.peekTok.TypeIs(token.SEMICOLON) && precedence < this.scanner.peekPrecedence() {
		infix := this.infixDecoders[this.scanner.peekTok.Type]
		if nil == infix {
			return leftExpr
		}
		this.scanner.nextToken()
		leftExpr = infix(leftExpr)
	}
	return leftExpr
}

func (this *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Tok:  this.scanner.curTok,
		Op:   this.scanner.curTok.Literal,
		Left: left,
	}
	preced := this.scanner.curPrecedence()
	this.scanner.nextToken()
	expr.Right = this.parseExpression(preced)
	return expr
}

func (this *Parser) parseCallExpression(left ast.Expression) ast.Expression {
	expr := &ast.Call{Tok: this.scanner.curTok, Func: left}
	expr.Args = this.parseCallArgs()
	return expr
}

func (this *Parser) parseCallArgs() ast.ExpressionSlice {
	args := ast.ExpressionSlice{}
	if this.scanner.peekTok.TypeIs(token.RPAREN) {
		this.scanner.nextToken()
		return args
	}
	this.scanner.nextToken()
	args = append(args, this.parseExpression(PRECED_LOWEST))

	for this.scanner.peekTok.TypeIs(token.COMMA) {
		this.scanner.nextToken()
		this.scanner.nextToken()
		args = append(args, this.parseExpression(PRECED_LOWEST))
	}
	if !this.scanner.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}
