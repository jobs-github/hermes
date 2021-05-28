package object

import (
	"Q/token"
	"fmt"
)

// Null : implement Object
type Null struct{}

func (this *Null) Type() ObjectType {
	return ObjectTypeNull
}

func (this *Null) Inspect() string {
	return "null"
}

func (this *Null) Opposite() (Object, error) {
	return nil, fmt.Errorf("Null.Opposite -> unsupported")
}

func (this *Null) Not() (Object, error) {
	return True, nil
}

func (this *Null) Calc(op *token.Token, right Object) (Object, error) {
	return right.calcNull(op, this)
}

func (this *Null) Call(args []Object, insideLoop bool) (Object, error) {
	return nil, fmt.Errorf("Null.Call -> unsupported")
}

func (this *Null) True() bool {
	return false
}

func (this *Null) Return() (bool, Object) {
	return false, nil
}

func (this *Null) Break() (bool, int) {
	return false, 0
}

func (this *Null) andInteger(left *Integer) Object {
	if 0 == left.Value {
		return left
	}
	return Nil
}

func (this *Null) andBoolean(left *Boolean) Object {
	if false == left.Value {
		return left
	}
	return Nil
}

func (this *Null) orInteger(left *Integer) Object {
	if 0 != left.Value {
		return left
	}
	return Nil
}

func (this *Null) orBoolean(left *Boolean) Object {
	if false != left.Value {
		return left
	}
	return Nil
}

func (this *Null) calcInteger(op *token.Token, left *Integer) (Object, error) {
	switch op.Type {
	case token.LT:
		return ToBoolean(false), nil
	case token.LEQ:
		return ToBoolean(false), nil
	case token.GT:
		return ToBoolean(true), nil
	case token.GEQ:
		return ToBoolean(true), nil
	case token.EQ:
		return ToBoolean(false), nil
	case token.NEQ:
		return ToBoolean(true), nil
	case token.AND:
		return this.andInteger(left), nil
	case token.OR:
		return this.orInteger(left), nil
	default:
		return nil, fmt.Errorf("Null.calcInteger -> unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Null) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	switch op.Type {
	case token.AND:
		return this.andBoolean(left), nil
	case token.OR:
		return this.orBoolean(left), nil
	default:
		return this.calcInteger(op, toInteger(left.Value))
	}
}

func (this *Null) calcNull(op *token.Token, left *Null) (Object, error) {
	switch op.Type {
	case token.LT:
		return ToBoolean(false), nil
	case token.LEQ:
		return ToBoolean(true), nil
	case token.GT:
		return ToBoolean(false), nil
	case token.GEQ:
		return ToBoolean(true), nil
	case token.EQ:
		return ToBoolean(true), nil
	case token.NEQ:
		return ToBoolean(false), nil
	case token.AND:
		return this, nil
	case token.OR:
		return this, nil
	default:
		return nil, fmt.Errorf("Null.calcNull -> unsupported op %v(%v)", op.Literal, op.Type)
	}
}
