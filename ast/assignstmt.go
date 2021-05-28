package ast

import (
	"Q/object"
	"bytes"
	"fmt"
)

// AssignStmt : implement Statement
type AssignStmt struct {
	Name  *Identifier
	Value Expression
}

func (this *AssignStmt) statementNode() {}
func (this *AssignStmt) TokenLiteral() string {
	return ""
}
func (this *AssignStmt) String() string {
	var out bytes.Buffer
	out.WriteString(this.Name.String())
	out.WriteString(" = ")
	if nil != this.Value {
		out.WriteString(this.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
func (this *AssignStmt) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	val, err := this.Value.Eval(env, insideLoop)
	if nil != err {
		return nil, fmt.Errorf("AssignStmt.Eval -> eval value | %v", err)
	}
	if err := env.Assign(this.Name.Value, val); nil != err {
		return nil, fmt.Errorf("AssignStmt.Eval -> env.Assign | %v", err)
	}
	return val, nil
}
