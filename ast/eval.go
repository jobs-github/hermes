package ast

import (
	"hermes/object"
)

func evalStatements(stmts StatementSlice) object.Object {
	var result object.Object
	for _, stmt := range stmts {
		result = stmt.Eval()
	}
	return result
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return right.Not()
	case "-":
		return right.Opposite()
	default:
		return nil
	}
}
