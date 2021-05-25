package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

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
func (this *VarStmt) Eval(env *object.Env) (object.Object, error) {
	val, err := this.Value.Eval(env)
	if nil != err {
		return nil, fmt.Errorf("VarStmt.Eval | %v", err)
	}
	env.Set(this.Name.Value, val)
	return val, nil
}
