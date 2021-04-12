package object

import (
	"Q/token"
)

type Object interface {
	Type() ObjectType
	Inspect() string
	Not() (Object, error)
	Opposite() (Object, error)
	Calc(op *token.Token, right Object) (Object, error)
	True() bool
	Return() (bool, Object)

	calcInteger(op *token.Token, left *Integer) (Object, error)
	calcBoolean(op *token.Token, left *Boolean) (Object, error)
	calcNull(op *token.Token, left *Null) (Object, error)
}
