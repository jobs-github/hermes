package object

import (
	"Q/token"
	"fmt"
)

const (
	ObjectTypeInteger ObjectType = iota
	ObjectTypeBoolean
	ObjectTypeNull
	ObjectTypeReturnValue
	ObjectTypeFunction
)

var (
	Nil   = &Null{}
	True  = &Boolean{Value: true}
	False = &Boolean{Value: false}
)

var (
	objectTypeStrings = map[ObjectType]string{
		ObjectTypeInteger:     "integer",
		ObjectTypeBoolean:     "boolean",
		ObjectTypeNull:        "null",
		ObjectTypeReturnValue: "return_value",
		ObjectTypeFunction:    "function",
	}
)

type ObjectType uint8

func ToString(t ObjectType) string {
	s, ok := objectTypeStrings[t]
	if ok {
		return s
	}
	return "undefined type"
}

func ToBoolean(v bool) *Boolean {
	if v {
		return True
	} else {
		return False
	}
}

func toInt64(v bool) int64 {
	if v {
		return 1
	} else {
		return 0
	}
}

func toInteger(v bool) *Integer {
	return &Integer{Value: toInt64(v)}
}

func infixNull(op *token.Token, right Object, method string) (Object, error) {
	switch op.Type {
	case token.LT:
		return ToBoolean(true), nil
	case token.LEQ:
		return ToBoolean(true), nil
	case token.GT:
		return ToBoolean(false), nil
	case token.GEQ:
		return ToBoolean(false), nil
	case token.EQ:
		return ToBoolean(false), nil
	case token.NEQ:
		return ToBoolean(true), nil
	case token.AND:
		return Nil, nil
	case token.OR:
		return right, nil
	default:
		return nil, fmt.Errorf("infixNull -> (%v) unsupported op %v(%v)", method, op.Literal, op.Type)
	}
}
