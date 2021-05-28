package ast

import (
	"Q/function"
	"Q/object"
	"Q/token"
	"bytes"
	"strings"
)

// Function : implement Expression
type Function struct {
	Tok  *token.Token
	Args IdentifierSlice
	Body *BlockStmt
}

func (this *Function) expressionNode() {}
func (this *Function) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *Function) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range this.Args {
		args = append(args, p.String())
	}
	out.WriteString(this.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	out.WriteString(this.Body.String())

	return out.String()
}
func (this *Function) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	return &object.Function{
		Fn: function.Function{
			Inspect:    this.inspect,
			ArgumentOf: this.argumentOf,
			Body:       this.body,
		},
		Args:     this.Args.values(),
		EvalBody: this.evalBody,
		Env:      env,
	}, nil
}

func (this *Function) evalBody(env *object.Env, insideLoop bool) (object.Object, error) {
	return this.Body.Eval(env, insideLoop)
}

func (this *Function) inspect() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range this.Args {
		args = append(args, p.String())
	}
	out.WriteString(this.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(") {\n")
	out.WriteString(this.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (this *Function) argumentOf(idx int) string {
	return this.Args[idx].String()
}

func (this *Function) body() string {
	return this.Body.String()
}
