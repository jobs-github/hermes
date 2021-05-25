package ast

import (
	"Q/object"
	"bytes"
)

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

func (this *Program) Eval(env *object.Env) (object.Object, error) {
	return evalStatements(env, this.Stmts, false)
}
