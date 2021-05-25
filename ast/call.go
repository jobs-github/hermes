package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
	"strings"
)

// Call : implement Expression
type Call struct {
	Tok  *token.Token
	Func Expression
	Args ExpressionSlice
}

func (this *Call) expressionNode() {}
func (this *Call) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Call) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range this.Args {
		args = append(args, a.String())
	}

	out.WriteString(this.Func.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (this *Call) Eval(env *object.Env) (object.Object, error) {
	fn, err := this.Func.Eval(env)
	if nil != err {
		return nil, fmt.Errorf("Call.Eval | %v", err)
	}

	args, err := this.Args.evalArgs(env)
	if nil != err {
		return nil, fmt.Errorf("Call.Eval | %v", err)
	}
	return fn.Call(args)
}
