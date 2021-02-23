package ast

import (
	"bytes"
	"hermes/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (this *Program) String() string {
	var out bytes.Buffer
	for _, s := range this.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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
func (this *Identifier) String() string {
	return this.Value
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
func (this *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(this.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(this.Name.String())
	out.WriteString(" = ")
	if nil != this.Value {
		out.WriteString(this.Value.String())
	}
	out.WriteString(";")
	return out.String()
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
func (this *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(this.TokenLiteral())
	out.WriteString(" ")

	if nil != this.ReturnValue {
		out.WriteString(this.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Tok  *token.Token
	Expr Expression
}

func (this *ExpressionStatement) statementNode() {}
func (this *ExpressionStatement) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *ExpressionStatement) String() string {
	if this.Expr != nil {
		return this.Expr.String()
	}
	return ""
}
