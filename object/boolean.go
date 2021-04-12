package object

import (
	"Q/token"
	"fmt"
)

// Boolean : implement Object
type Boolean struct {
	Value bool
}

func (this *Boolean) Type() ObjectType {
	return ObjectTypeBoolean
}

func (this *Boolean) Inspect() string {
	return fmt.Sprintf("%v", this.Value)
}

func (this *Boolean) Opposite() (Object, error) {
	if this.Value {
		return &Integer{Value: -1}, nil
	} else {
		return &Integer{Value: 0}, nil
	}
}

func (this *Boolean) Not() (Object, error) {
	if this.Value {
		return False, nil
	} else {
		return True, nil
	}
}

func (this *Boolean) Calc(op *token.Token, right Object) (Object, error) {
	return right.calcBoolean(op, this)
}

func (this *Boolean) True() bool {
	return this.Value
}

func (this *Boolean) Return() (bool, Object) {
	return false, nil
}

func (this *Boolean) calcInteger(op *token.Token, left *Integer) (Object, error) {
	right := toInteger(this.Value)
	return right.calcInteger(op, left)
}

func (this *Boolean) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	switch op.Type {
	case token.ADD:
		return &Integer{Value: toInt64(left.Value) + toInt64(this.Value)}, nil
	case token.SUB:
		return &Integer{Value: toInt64(left.Value) - toInt64(this.Value)}, nil
	case token.MUL:
		return &Integer{Value: toInt64(left.Value) * toInt64(this.Value)}, nil
	case token.DIV:
		return &Integer{Value: toInt64(left.Value) / toInt64(this.Value)}, nil
	case token.MOD:
		return &Integer{Value: toInt64(left.Value) % toInt64(this.Value)}, nil
	case token.LT:
		return ToBoolean(toInt64(left.Value) < toInt64(this.Value)), nil
	case token.LEQ:
		return ToBoolean(toInt64(left.Value) <= toInt64(this.Value)), nil
	case token.GT:
		return ToBoolean(toInt64(left.Value) > toInt64(this.Value)), nil
	case token.GEQ:
		return ToBoolean(toInt64(left.Value) >= toInt64(this.Value)), nil
	case token.EQ:
		return ToBoolean(left.Value == this.Value), nil
	case token.NEQ:
		return ToBoolean(left.Value != this.Value), nil
	case token.AND:
		return ToBoolean(left.Value && this.Value), nil
	case token.OR:
		return ToBoolean(left.Value || this.Value), nil
	default:
		return nil, fmt.Errorf("Boolean.calcBoolean: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Boolean) calcNull(op *token.Token, left *Null) (Object, error) {
	return infixNull(op, this, "Boolean.calcNull")
}
