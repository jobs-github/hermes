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

// Program : implement Node
type Program struct {
	Stmts StatementSlice
}

func (this *Program) TokenLiteral() string {
	if len(this.Stmts) > 0 {
		return this.Stmts[0].TokenLiteral()
	}
	return ""
}

func (this *Program) String() string {
	var out bytes.Buffer
	for _, s := range this.Stmts {
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

// VarStmt : implement Statement
type VarStmt struct {
	Tok   *token.Token
	Name  *Identifier
	Value Expression
}

func (this *VarStmt) statementNode() {}
func (this *VarStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *VarStmt) String() string {
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

// ReturnStmt : implement Statement
type ReturnStmt struct {
	Tok         *token.Token
	ReturnValue Expression
}

func (this *ReturnStmt) statementNode() {}
func (this *ReturnStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *ReturnStmt) String() string {
	var out bytes.Buffer
	out.WriteString(this.TokenLiteral())
	out.WriteString(" ")

	if nil != this.ReturnValue {
		out.WriteString(this.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStmt : implement Statement
type ExpressionStmt struct {
	Tok  *token.Token
	Expr Expression
}

func (this *ExpressionStmt) statementNode() {}
func (this *ExpressionStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *ExpressionStmt) String() string {
	if this.Expr != nil {
		return this.Expr.String()
	}
	return ""
}

type BlockStmt struct {
	Tok   *token.Token // {
	Stmts StatementSlice
}

func (this *BlockStmt) statementNode() {}
func (this *BlockStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *BlockStmt) String() string {
	var out bytes.Buffer
	for _, s := range this.Stmts {
		out.WriteString(s.String())
	}
	return out.String()
}

// Integer : implement Expression
type Integer struct {
	Tok   *token.Token
	Value int64
}

func (this *Integer) expressionNode() {}
func (this *Integer) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Integer) String() string {
	return this.Tok.Literal
}

// Boolean : implement Expression
type Boolean struct {
	Tok   *token.Token
	Value bool
}

func (this *Boolean) expressionNode() {}
func (this *Boolean) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Boolean) String() string {
	return this.Tok.Literal
}

type IfClause struct {
	If   Expression
	Then *BlockStmt
}

type IfClauseSlice []*IfClause

type IfExpression struct {
	Tok     *token.Token
	Clauses IfClauseSlice
	Else    *BlockStmt
}

func (this *IfExpression) expressionNode() {}
func (this *IfExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	for _, clause := range this.Clauses {
		out.WriteString(clause.If.String())
		out.WriteString(" ")
		out.WriteString(clause.Then.String())
	}
	if nil != this.Else {
		out.WriteString("else ")
		out.WriteString(this.Else.String())
	}
	return out.String()
}

// PrefixExpression : implement Expression
type PrefixExpression struct {
	Tok   *token.Token
	Op    string
	Right Expression
}

func (this *PrefixExpression) expressionNode() {}
func (this *PrefixExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Op)
	out.WriteString(this.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression : implement Expression
type InfixExpression struct {
	Tok   *token.Token
	Left  Expression
	Op    string
	Right Expression
}

func (this *InfixExpression) expressionNode() {}
func (this *InfixExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Left.String())
	out.WriteString(" ")
	out.WriteString(this.Op)
	out.WriteString(" ")
	out.WriteString(this.Right.String())
	out.WriteString(")")
	return out.String()
}
