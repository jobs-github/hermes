package ast

import (
	"Q/object"
	"Q/token"
	"fmt"
)

func evalStatements(env *object.Env, stmts StatementSlice, blockStmts bool) (object.Object, error) {
	var result object.Object
	for _, stmt := range stmts {
		if v, err := stmt.Eval(env); nil != err {
			return nil, fmt.Errorf("evalStatements error: %v", err)
		} else {
			if needReturn, returnValue := v.Return(); needReturn {
				if blockStmts {
					// it stops execution in a possible deeper block statement and bubbles up to Program.Eval
					// where it finally get's unwrapped
					return v, nil
				} else {
					return returnValue, nil
				}
			}
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
