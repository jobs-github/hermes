package ast

import (
	"Q/object"
	"Q/token"
	"bytes"
)

// BlockStmt : implement Statement
type BlockStmt struct {
	Tok   *token.Token // {
	Stmts StatementSlice
}

func (this *BlockStmt) statementNode() {}
func (this *BlockStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *BlockStmt) String() string {
	var out bytes.Buffer
	for _, s := range this.Stmts {
		out.WriteString(s.String())
	}
	return out.String()
}
func (this *BlockStmt) Eval(env *object.Env) (object.Object, error) {
	return evalStatements(env, this.Stmts, true)
}
