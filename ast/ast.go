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
	Tok   *token.Token
	Value string
}

func (this *Identifier) expressionNode() {}
func (this *Identifier) TokenLiteral() string {
	return this.Tok.Literal
}

// VarStatement : implement Expression
type VarStatement struct {
	Tok   *token.Token
	Name  *Identifier
	Value Expression
}

func (this *VarStatement) statementNode() {}
func (this *VarStatement) TokenLiteral() string {
	return this.Tok.Literal
}
