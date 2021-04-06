package parser

import (
	"Q/ast"
	"Q/token"
	"fmt"
	"strconv"
)

type tokenDecoder interface {
	decode() ast.Expression
}

type tokenDecoderMap map[token.TokenType]tokenDecoder

func newTokenDecoders(
	s *scanner,
	parseExpression parseExpressionFn,
	parseBlockStmt parseBlockStmtFn,
) tokenDecoderMap {
	identifierDecoder := &identifier{s}
	integerDecoder := &integer{s}
	booleanDecoder := &boolean{s}
	prefixExprDecoder := &prefixExpr{s, parseExpression}
	groupedExprDecoder := &groupedExpr{s, parseExpression}
	ifExprDecoder := &ifExpr{s, parseExpression, parseBlockStmt}
	funcDecoder := &funcLiteral{s, parseExpression, parseBlockStmt}

	return tokenDecoderMap{
		token.IDENT:  identifierDecoder,
		token.INT:    integerDecoder,
		token.TRUE:   booleanDecoder,
		token.FALSE:  booleanDecoder,
		token.NOT:    prefixExprDecoder,
		token.SUB:    prefixExprDecoder,
		token.LPAREN: groupedExprDecoder,
		token.IF:     ifExprDecoder,
		token.FUNC:   funcDecoder,
	}
}

// identifier : implement tokenDecoder
type identifier struct {
	scanner *scanner
}

func (this *identifier) decode() ast.Expression {
	return &ast.Identifier{Tok: this.scanner.curTok, Value: this.scanner.curTok.Literal}
}

// boolean : implement tokenDecoder
type boolean struct {
	scanner *scanner
}

func (this *boolean) decode() ast.Expression {
	return &ast.Boolean{Tok: this.scanner.curTok, Value: this.scanner.curTok.TypeIs(token.TRUE)}
}

// integer : implement tokenDecoder
type integer struct {
	scanner *scanner
}

func (this *integer) decode() ast.Expression {
	expr := &ast.Integer{Tok: this.scanner.curTok}
	val, err := strconv.ParseInt(this.scanner.curTok.Literal, 0, 64)
	if nil != err {
		this.scanner.appendError(fmt.Sprintf("could not parse %v as integer", this.scanner.curTok.Literal))
		return nil
	}
	expr.Value = val
	return expr
}

// prefixExpr : implement tokenDecoder
type prefixExpr struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *prefixExpr) decode() ast.Expression {
	expr := &ast.PrefixExpression{
		Tok: this.scanner.curTok,
		Op:  this.scanner.curTok,
	}
	this.scanner.nextToken()
	expr.Right = this.parseExpression(PRECED_PREFIX)
	return expr
}

// groupedExpr : implement tokenDecoder
type groupedExpr struct {
	scanner         *scanner
	parseExpression parseExpressionFn
}

func (this *groupedExpr) decode() ast.Expression {
	this.scanner.nextToken()
	expr := this.parseExpression(PRECED_LOWEST)
	if !this.scanner.expectPeek(token.RPAREN) {
		return nil
	}
	return expr
}

// ifExpr : implement tokenDecoder
type ifExpr struct {
	scanner         *scanner
	parseExpression parseExpressionFn
	parseBlockStmt  parseBlockStmtFn
}

func (this *ifExpr) decodeClause() *ast.IfClause {
	clause := &ast.IfClause{}
	if !this.scanner.expectPeek(token.LPAREN) {
		return nil
	}
	this.scanner.nextToken()
	clause.If = this.parseExpression(PRECED_LOWEST)
	if !this.scanner.expectPeek(token.RPAREN) {
		return nil
	}
	if !this.scanner.expectPeek(token.LBRACE) {
		return nil
	}
	clause.Then = this.parseBlockStmt()
	return clause
}

func (this *ifExpr) decodeElseIf(expr *ast.IfExpression) bool {
	this.scanner.nextToken()
	this.scanner.nextToken()

	clause := this.decodeClause()
	if nil == clause {
		return false
	}
	expr.Clauses = append(expr.Clauses, clause)
	return true
}

func (this *ifExpr) decodeElse(expr *ast.IfExpression) bool {
	this.scanner.nextToken()
	if !this.scanner.expectPeek(token.LBRACE) {
		return false
	}
	expr.Else = this.parseBlockStmt()
	return true
}

func (this *ifExpr) decode() ast.Expression {
	expr := &ast.IfExpression{Tok: this.scanner.curTok, Clauses: ast.IfClauseSlice{}}

	clause := this.decodeClause()
	if nil == clause {
		return nil
	}
	expr.Clauses = append(expr.Clauses, clause)
	for {
		if this.scanner.expectPeek2(token.ELSE, token.IF) {
			if !this.decodeElseIf(expr) {
				return nil
			}
		} else if this.scanner.peekTok.TypeIs(token.ELSE) {
			if !this.decodeElse(expr) {
				return nil
			}
			break
		} else {
			break
		}
	}
	return expr
}

// funcLiteral : implement tokenDecoder
type funcLiteral struct {
	scanner         *scanner
	parseExpression parseExpressionFn
	parseBlockStmt  parseBlockStmtFn
}

func (this *funcLiteral) parseArgs() ast.IdentifierSlice {
	args := ast.IdentifierSlice{}
	if this.scanner.peekTok.TypeIs(token.RPAREN) {
		this.scanner.nextToken()
		return args
	}
	this.scanner.nextToken()
	ident := &ast.Identifier{Tok: this.scanner.curTok, Value: this.scanner.curTok.Literal}
	args = append(args, ident)

	for this.scanner.peekTok.TypeIs(token.COMMA) {
		this.scanner.nextToken()
		this.scanner.nextToken()
		ident := &ast.Identifier{Tok: this.scanner.curTok, Value: this.scanner.curTok.Literal}
		args = append(args, ident)
	}

	if !this.scanner.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}

func (this *funcLiteral) decode() ast.Expression {
	lit := &ast.Function{Tok: this.scanner.curTok}
	if !this.scanner.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Args = this.parseArgs()
	if !this.scanner.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = this.parseBlockStmt()
	return lit
}
