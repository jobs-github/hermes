package object

import (
	"Q/token"
	"fmt"
)

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
	return nil, fmt.Errorf("ReturnValue.Opposite -> unsupported")
}

func (this *ReturnValue) Not() (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Opposite -> unsupported")
}

func (this *ReturnValue) Calc(op *token.Token, right Object) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Opposite -> unsupported")
}

func (this *ReturnValue) Call(args []Object, insideLoop bool) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.Call -> unsupported")
}

func (this *ReturnValue) True() bool {
	return false
}

func (this *ReturnValue) Return() (bool, Object) {
	return true, this.Value
}

func (this *ReturnValue) Break() (bool, int) {
	return false, 0
}

func (this *ReturnValue) calcInteger(op *token.Token, left *Integer) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcInteger -> unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *ReturnValue) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcBoolean -> unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *ReturnValue) calcNull(op *token.Token, left *Null) (Object, error) {
	return nil, fmt.Errorf("ReturnValue.calcNull -> unsupported op %v(%v)", op.Literal, op.Type)
}
