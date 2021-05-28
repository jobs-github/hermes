package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
	"fmt"
)

// ForExpression : implement Expression
type ForExpression struct {
	Tok  *token.Token
	Loop *BlockStmt
}

func (this *ForExpression) expressionNode() {}
func (this *ForExpression) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *ForExpression) String() string {
	var out bytes.Buffer
	out.WriteString("for {")
	if nil != this.Loop {
		out.WriteString(this.Loop.String())
	}
	out.WriteString("}")
	return out.String()
}
func (this *ForExpression) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	var rc object.Object
	for {
		v, err := this.Loop.Eval(env, true)
		if nil != err {
			return nil, fmt.Errorf("ForExpression.Eval | %v", err)
		}
		if isBreak, _ := v.Break(); isBreak {
			rc = v
			break
		}
		if needReturn, _ := v.Return(); needReturn {
			return v, nil
		}
	}
	return rc, nil
}
