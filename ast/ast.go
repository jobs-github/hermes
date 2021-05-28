package ast

import (
	"Q/object"
	"Q/token"
	"fmt"
)

type Node interface {
	TokenLiteral() string
	String() string
	Eval(env *object.Env, insideLoop bool) (object.Object, error)
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type ExpressionSlice []Expression

func (this *ExpressionSlice) evalArgs(env *object.Env, insideLoop bool) ([]object.Object, error) {
	result := []object.Object{}
	for _, expr := range *this {
		evaluated, err := expr.Eval(env, insideLoop)
		if nil != err {
			return nil, fmt.Errorf("ExpressionSlice.eval | %v", err)
		}
		result = append(result, evaluated)
	}
	return result, nil
}

type StatementSlice []Statement

func (this *StatementSlice) eval(isBlockStmts bool, env *object.Env, insideLoop bool) (object.Object, error) {
	var result object.Object
	for _, stmt := range *this {
		if v, err := stmt.Eval(env, insideLoop); nil != err {
			return nil, fmt.Errorf("evalStatements | %v", err)
		} else {
			if needReturn, returnValue := v.Return(); needReturn {
				if isBlockStmts {
					// it stops execution in a possible deeper block statement and bubbles up to Program.Eval
					// where it finally get's unwrapped
					return v, nil
				} else {
					return returnValue, nil
				}
			}
			if insideLoop {
				isBreak, _ := v.Break()
				if isBreak {
					return v, nil
				}
			} else { // outside loop
				isBreak, breakCount := v.Break()
				if isBreak && 1 == breakCount { // orginal break
					return nil, fmt.Errorf("evalStatements -> 'break' outside loop")
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
		return nil, fmt.Errorf("evalPrefixExpression -> unsupport op %v(%v)", op.Literal, op.Type)
	}
}
