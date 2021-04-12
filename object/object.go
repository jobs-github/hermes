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
		return nil, fmt.Errorf("infixNull: (%v) unsupported op %v(%v)", method, op.Literal, op.Type)
	}
}

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

func (this *Integer) Return() (bool, Object) {
	return false, nil
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

func (this *Null) Return() (bool, Object) {
	return false, nil
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
		return nil, fmt.Errorf("Null.calcInteger: unsupported op %v(%v)", op.Literal, op.Type)
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
		return nil, fmt.Errorf("Null.calcNull: unsupported op %v(%v)", op.Literal, op.Type)
	}
}

// ReturnValue : implement Object
type ReturnValue struct {
	Value Object
}

func (this *ReturnValue) Type() ObjectType {
	return ObjectTypeReturnValue
}

func (this *ReturnValue) Inspect() string {
	return this.Value.Inspect()
}

func (this *ReturnValue) Opposite() (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Opposite: not supported")
}

func (this *ReturnValue) Not() (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Opposite: not supported")
}

func (this *ReturnValue) Calc(op *token.Token, right Object) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Opposite: not supported")
}

func (this *ReturnValue) True() bool {
	return false
}

func (this *ReturnValue) Return() (bool, Object) {
	return true, this.Value
}

func (this *ReturnValue) calcInteger(op *token.Token, left *Integer) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcInteger: unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *ReturnValue) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcBoolean: unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *ReturnValue) calcNull(op *token.Token, left *Null) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcNull: unsupported op %v(%v)", op.Literal, op.Type)
}
