package object

import (
	"Q/token"
	"fmt"
)

const (
	ObjectTypeInteger ObjectType = iota
	ObjectTypeBoolean
	ObjectTypeNull
)

var (
	Nil   = &Null{}
	True  = &Boolean{Value: true}
	False = &Boolean{Value: false}
)

var (
	objectTypeStrings = map[ObjectType]string{
		ObjectTypeInteger: "integer",
		ObjectTypeBoolean: "boolean",
		ObjectTypeNull:    "null",
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

type Object interface {
	Type() ObjectType
	Inspect() string
	Not() (Object, error)
	Opposite() (Object, error)
	Calc(op *token.Token, right Object) (Object, error)
	True() bool

	calcInteger(op *token.Token, left *Integer) (Object, error)
	calcBoolean(op *token.Token, left *Boolean) (Object, error)
	calcNull(op *token.Token, left *Null) (Object, error)
}

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

func (this *Integer) True() bool {
	if 0 == this.Value {
		return false
	}
	return true
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
		return nil, fmt.Errorf("Integer.calcInteger: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Integer) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Integer.calcBoolean: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Integer) calcNull(op *token.Token, left *Null) (Object, error) {
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Integer.calcNull: unsupported op %v(%v)", op.Literal, op.Type)
	}
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

func (this *Boolean) calcInteger(op *token.Token, left *Integer) (Object, error) {
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Boolean.calcInteger: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Boolean) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	switch op.Type {
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
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Boolean.calcNull: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

// Null : implement Object
type Null struct{}

func (this *Null) Type() ObjectType {
	return ObjectTypeNull
}

func (this *Null) Inspect() string {
	return "null"
}

func (this *Null) Opposite() (Object, error) {
	return nil, fmt.Errorf("Null.Opposite: not supported")
}

func (this *Null) Not() (Object, error) {
	return True, nil
}

func (this *Null) Calc(op *token.Token, right Object) (Object, error) {
	return right.calcNull(op, this)
}

func (this *Null) True() bool {
	return false
}

func (this *Null) calcInteger(op *token.Token, left *Integer) (Object, error) {
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Null.calcInteger: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Null) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	switch op.Type {
	// TODO
	default:
		return nil, fmt.Errorf("Null.calcBoolean: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

func (this *Null) calcNull(op *token.Token, left *Null) (Object, error) {
	switch op.Type {
	case token.EQ:
		return ToBoolean(true), nil
	case token.NEQ:
		return ToBoolean(false), nil
	default:
		return nil, fmt.Errorf("Null.calcNull: unsupported op %v(%v)", op.Literal, op.Type)
	}
}
