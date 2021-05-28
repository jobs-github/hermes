package object

import (
	"Q/token"
	"fmt"
)

// Integer : implement Object
type Integer struct {
	Value int64
}

func (this *Integer) Type() ObjectType {
	return ObjectTypeInteger
}

func (this *Integer) Inspect() string {
	return fmt.Sprintf("%v", this.Value)
}

func (this *Integer) Opposite() (Object, error) {
	return &Integer{Value: -this.Value}, nil
}

func (this *Integer) Not() (Object, error) {
	if 0 == this.Value {
		return True, nil
	} else {
		return False, nil
	}
}

func (this *Integer) Calc(op *token.Token, right Object) (Object, error) {
	return right.calcInteger(op, this)
}

func (this *Integer) Call(args []Object, insideLoop bool) (Object, error) {
	return nil, fmt.Errorf("Integer.Call -> unsupported")
}

func (this *Integer) True() bool {
	if 0 == this.Value {
		return false
	}
	return true
}

func (this *Integer) Return() (bool, Object) {
	return false, nil
}

func (this *Integer) Break() (bool, int) {
	return false, 0
}

func (this *Integer) calcInteger(op *token.Token, left *Integer) (Object, error) {
	switch op.Type {
	case token.ADD:
		return &Integer{Value: left.Value + this.Value}, nil
	case token.SUB:
		return &Integer{Value: left.Value - this.Value}, nil
	case token.MUL:
		return &Integer{Value: left.Value * this.Value}, nil
	case token.DIV:
		return &Integer{Value: left.Value / this.Value}, nil
	case token.MOD:
		return &Integer{Value: left.Value % this.Value}, nil
	case token.LT:
		return ToBoolean(left.Value < this.Value), nil
	case token.LEQ:
		return ToBoolean(left.Value <= this.Value), nil
	case token.GT:
		return ToBoolean(left.Value > this.Value), nil
	case token.GEQ:
		return ToBoolean(left.Value >= this.Value), nil
	case token.EQ:
		return ToBoolean(left.Value == this.Value), nil
	case token.NEQ:
		return ToBoolean(left.Value != this.Value), nil
	case token.AND:
		return this.and(left)
	case token.OR:
		return this.or(left)
	default:
		return nil, fmt.Errorf("Integer.calcInteger -> unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Integer) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	return this.calcInteger(op, toInteger(left.Value))
}

func (this *Integer) calcNull(op *token.Token, left *Null) (Object, error) {
	return infixNull(op, this, "Integer.calcNull")
}

func (this *Integer) and(left *Integer) (Object, error) {
	if 0 == left.Value {
		return left, nil
	}
	return this, nil
}

func (this *Integer) or(left *Integer) (Object, error) {
	if 0 != left.Value {
		return left, nil
	}
	return this, nil
}
