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

// VarStatement : implement Statement
type VarStatement struct {
	Tok   *token.Token
	Name  *Identifier
	Value Expression
}

func (this *VarStatement) statementNode() {}
func (this *VarStatement) TokenLiteral() string {
	return this.Tok.Literal
}

// ReturnStatement : implement Statement
type ReturnStatement struct {
	Tok         *token.Token
	ReturnValue Expression
}

func (this *ReturnStatement) statementNode() {}
func (this *ReturnStatement) TokenLiteral() string {
	return this.Tok.Literal
}
