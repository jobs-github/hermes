package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
)

// BreakStmt : implement Statement
type BreakStmt struct {
	Tok *token.Token
}

func (this *BreakStmt) statementNode() {}
func (this *BreakStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *BreakStmt) String() string {
	var out bytes.Buffer
	out.WriteString(this.TokenLiteral())
	out.WriteString(";")
	return out.String()
}
func (this *BreakStmt) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	return object.NewBreak(), nil
}
