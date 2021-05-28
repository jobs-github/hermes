package ast

import (
	"Q/object"
	"Q/token"
)

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
func (this *Boolean) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	return object.ToBoolean(this.Value), nil
}
