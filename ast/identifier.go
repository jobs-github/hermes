package ast

import (
	"Q/object"
	"Q/token"
	"fmt"
)

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
func (this *Identifier) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	val, ok := env.Get(this.Value)
	if !ok {
		return nil, fmt.Errorf("Identifier.Eval -> `%v` not found", this.Value)
	}
	return val, nil
}

type IdentifierSlice []*Identifier

func (this *IdentifierSlice) values() []string {
	v := []string{}
	for _, i := range *this {
		v = append(v, i.Value)
	}
	return v
}
