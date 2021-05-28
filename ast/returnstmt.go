package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

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
func (this *ReturnStmt) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	val, err := this.ReturnValue.Eval(env, insideLoop)
	if nil != err {
		return nil, fmt.Errorf("ReturnStmt.Eval | %v", err)
	}
	return &object.ReturnValue{Value: val}, nil
}
