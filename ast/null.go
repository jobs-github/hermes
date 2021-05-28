package ast

import (
	"Q/object"
	"Q/token"
)

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
func (this *Null) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	return object.Nil, nil
}
