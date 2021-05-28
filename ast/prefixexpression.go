package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

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
func (this *PrefixExpression) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	right, err := this.Right.Eval(env, insideLoop)
	if nil != err {
		return nil, fmt.Errorf("PrefixExpression.Eval -> this.Right.Eval() error | %v", err)
	}
	return evalPrefixExpression(this.Op, right)
}
