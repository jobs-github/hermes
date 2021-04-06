package ast

import (
	"Q/object"
	"Q/token"
	"fmt"
)

func evalStatements(stmts StatementSlice) (object.Object, error) {
	var result object.Object
	for _, stmt := range stmts {
		if v, err := stmt.Eval(); nil != err {
			return nil, fmt.Errorf("evalStatements error: %v", err)
		} else {
			result = v
		}
	}
	return result, nil
}

func evalPrefixExpression(op *token.Token, right object.Object) (object.Object, error) {
	switch op.Type {
	case token.NOT:
		return right.Not()
	case token.SUB:
		return right.Opposite()
	default:
		return nil, fmt.Errorf("evalPrefixExpression: unsupport op %v(%v)", op.Literal, op.Type)
	}
}

func evalInfixExpression(op *token.Token, left object.Object, right object.Object) (object.Object, error) {
	return left.Calc(op, right)
}
