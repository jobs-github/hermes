package ast

import (
	"hermes/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type StatementSlice []Statement

type Program struct {
	Statements StatementSlice
}

func (this *Program) TokenLiteral() string {
	if len(this.Statements) > 0 {
		return this.Statements[0].TokenLiteral()
	}
	return ""
}

// Identifier : implement Expression
type Identifier struct {
	tok   *token.Token
	value string
}

func (this *Identifier) expressionNode() {}
func (this *Identifier) TokenLiteral() string {
	return this.tok.Literal
}

// VarStatement : implement Expression
type VarStatement struct {
	tok   *token.Token
	name  *Identifier
	value Expression
}

func (this *VarStatement) statementNode() {}
func (this *VarStatement) TokenLiteral() string {
	return this.tok.Literal
}
