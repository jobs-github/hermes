package ast

import (
	"Q/object"
	"Q/token"
)

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
func (this *Integer) Eval(env *object.Env) (object.Object, error) {
	return &object.Integer{Value: this.Value}, nil
}
