package ast

import (
	"hermes/object"
	"hermes/token"
)

func evalStatements(stmts StatementSlice) object.Object {
	var result object.Object
	for _, stmt := range stmts {
		result = stmt.Eval()
	}
	return result
}

func evalPrefixExpression(op *token.Token, right object.Object) object.Object {
	switch op.Type {
	case token.NOT:
		return right.Not()
	case token.SUB:
		return right.Opposite()
	default:
		return nil
	}
}

func evalInfixExpression(op *token.Token, left object.Object, right object.Object) object.Object {
	return left.Calc(op, right)
}
