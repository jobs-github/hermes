package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
	Eval() (object.Object, error)
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type ExpressionSlice []Expression

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

func (this *Program) Eval() (object.Object, error) {
	return evalStatements(this.Stmts)
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
func (this *Identifier) Eval() (object.Object, error) {
	// TODO
	return nil, fmt.Errorf("Identifier.Eval not implement")
}

type IdentifierSlice []*Identifier

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
func (this *VarStmt) Eval() (object.Object, error) {
	// TODO
	return nil, fmt.Errorf("VarStmt.Eval not implement")
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
func (this *ReturnStmt) Eval() (object.Object, error) {
	// TODO
	return nil, fmt.Errorf("ReturnStmt.Eval not implement")
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
func (this *ExpressionStmt) Eval() (object.Object, error) {
	return this.Expr.Eval()
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
func (this *BlockStmt) Eval() (object.Object, error) {
	return evalStatements(this.Stmts)
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
func (this *Integer) Eval() (object.Object, error) {
	return &object.Integer{Value: this.Value}, nil
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
func (this *Boolean) Eval() (object.Object, error) {
	return object.ToBoolean(this.Value), nil
}

// Null : implement Expression
type Null struct {
	Tok *token.Token
}

func (this *Null) expressionNode() {}
func (this *Null) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Null) String() string {
	return this.Tok.Literal
}
func (this *Null) Eval() (object.Object, error) {
	return object.Nil, nil
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

	for i, clause := range this.Clauses {
		if 0 == i {
			out.WriteString("if")
		} else {
			out.WriteString("else if")
		}
		out.WriteString(clause.If.String())
		out.WriteString("{")
		out.WriteString(clause.Then.String())
		out.WriteString("}")
	}
	if nil != this.Else {
		out.WriteString("else ")
		out.WriteString("{")
		out.WriteString(this.Else.String())
		out.WriteString("}")
	}
	return out.String()
}
func (this *IfExpression) Eval() (object.Object, error) {
	for _, clause := range this.Clauses {
		cond, err := clause.If.Eval()
		if nil != err {
			return nil, fmt.Errorf("IfExpression.Eval: %v, err: %v", clause.If.String(), err)
		}
		if cond.True() {
			return clause.Then.Eval()
		}
	}
	if nil != this.Else {
		return this.Else.Eval()
	}
	return object.Nil, nil
}

// Function : implement Expression
type Function struct {
	Tok  *token.Token
	Args IdentifierSlice
	Body *BlockStmt
}

func (this *Function) expressionNode() {}
func (this *Function) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Function) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range this.Args {
		args = append(args, p.String())
	}
	out.WriteString(this.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	out.WriteString(this.Body.String())

	return out.String()
}
func (this *Function) Eval() (object.Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.Eval not implement")
}

// Call : implement Expression
type Call struct {
	Tok  *token.Token
	Func Expression
	Args ExpressionSlice
}

func (this *Call) expressionNode() {}
func (this *Call) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Call) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range this.Args {
		args = append(args, a.String())
	}

	out.WriteString(this.Func.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (this *Call) Eval() (object.Object, error) {
	// TODO
	return nil, fmt.Errorf("Call.Eval not implement")
}

// PrefixExpression : implement Expression
type PrefixExpression struct {
	Tok   *token.Token
	Op    *token.Token
	Right Expression
}

func (this *PrefixExpression) expressionNode() {}
func (this *PrefixExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Op.Literal)
	out.WriteString(this.Right.String())
	out.WriteString(")")
	return out.String()
}
func (this *PrefixExpression) Eval() (object.Object, error) {
	right, err := this.Right.Eval()
	if nil != err {
		return nil, fmt.Errorf("PrefixExpression.Eval: this.Right.Eval() error, %v", err)
	}
	return evalPrefixExpression(this.Op, right)
}

// InfixExpression : implement Expression
type InfixExpression struct {
	Tok   *token.Token
	Left  Expression
	Op    *token.Token
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
	out.WriteString(this.Op.Literal)
	out.WriteString(" ")
	out.WriteString(this.Right.String())
	out.WriteString(")")
	return out.String()
}
func (this *InfixExpression) Eval() (object.Object, error) {
	left, err := this.Left.Eval()
	if nil != err {
		return nil, fmt.Errorf("InfixExpression.Eval: this.Left.Eval() error, %v", err)
	}
	right, err := this.Right.Eval()
	if nil != err {
		return nil, fmt.Errorf("InfixExpression.Eval: this.Right.Eval() error, %v", err)
	}
	return left.Calc(this.Op, right)
}
