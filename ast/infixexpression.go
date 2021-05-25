package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

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
func (this *InfixExpression) Eval(env *object.Env) (object.Object, error) {
	left, err := this.Left.Eval(env)
	if nil != err {
		return nil, fmt.Errorf("InfixExpression.Eval -> this.Left.Eval() error | %v", err)
	}
	right, err := this.Right.Eval(env)
	if nil != err {
		return nil, fmt.Errorf("InfixExpression.Eval -> this.Right.Eval() error | %v", err)
	}
	return left.Calc(this.Op, right)
}
