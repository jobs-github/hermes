package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

type IfClause struct {
	If   Expression
	Then *BlockStmt
}

type IfClauseSlice []*IfClause

// IfExpression : implement Expression
type IfExpression struct {
	Tok     *token.Token
	Clauses IfClauseSlice
	Else    *BlockStmt
}

func (this *IfExpression) expressionNode() {}
func (this *IfExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *IfExpression) String() string {
	var out bytes.Buffer

	for i, clause := range this.Clauses {
		if 0 == i {
			out.WriteString("if")
		} else {
			out.WriteString("else if")
		}
		out.WriteString(clause.If.String())
		out.WriteString("{")
		out.WriteString(clause.Then.String())
		out.WriteString("}")
	}
	if nil != this.Else {
		out.WriteString("else ")
		out.WriteString("{")
		out.WriteString(this.Else.String())
		out.WriteString("}")
	}
	return out.String()
}
func (this *IfExpression) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	for _, clause := range this.Clauses {
		cond, err := clause.If.Eval(env, insideLoop)
		if nil != err {
			return nil, fmt.Errorf("IfExpression.Eval -> %v | %v", clause.If.String(), err)
		}
		if cond.True() {
			return clause.Then.Eval(env, insideLoop)
		}
	}
	if nil != this.Else {
		return this.Else.Eval(env, insideLoop)
	}
	return object.Nil, nil
}
