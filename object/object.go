package object

import (
	"fmt"
	"hermes/token"
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
	CalcInteger(op *token.Token, left *Integer) (Object, error)
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
	return right.CalcInteger(op, this)
}

func (this *Integer) CalcInteger(op *token.Token, left *Integer) (Object, error) {
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
	case token.GT:
		return ToBoolean(left.Value > this.Value), nil
	case token.EQ:
		return ToBoolean(left.Value == this.Value), nil
	case token.NEQ:
		return ToBoolean(left.Value != this.Value), nil
	// TODO
	default:
		return nil, fmt.Errorf("Integer.CalcInteger: unsupported op %v(%v)", op.Literal, op.Type)
	}
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
	// TODO
	return nil, fmt.Errorf("Boolean.Calc: not implement")
}

func (this *Boolean) CalcInteger(op *token.Token, left *Integer) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Boolean.CalcInteger: not implement")
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
	// TODO
	return nil, fmt.Errorf("Null.Calc: not implement")
}

func (this *Null) CalcInteger(op *token.Token, left *Integer) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Null.CalcInteger: not implement")
}
