package object

import "fmt"

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
	Not() Object
	Opposite() Object
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

func (this *Integer) Opposite() Object {
	return &Integer{Value: -this.Value}
}

func (this *Integer) Not() Object {
	if 0 == this.Value {
		return True
	} else {
		return False
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

func (this *Boolean) Opposite() Object {
	return nil
}

func (this *Boolean) Not() Object {
	if this.Value {
		return False
	} else {
		return True
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

func (this *Null) Opposite() Object {
	return nil
}

func (this *Null) Not() Object {
	return True
}
