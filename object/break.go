package object

import (
	"Q/token"
	"fmt"
)

func NewBreak() Object {
	return &BreakObject{count: 0}
}

// BreakObject : implement Object
type BreakObject struct {
	count int
}

func (this *BreakObject) Type() ObjectType {
	return ObjectTypeBreakObject
}

func (this *BreakObject) Inspect() string {
	return ToString(ObjectTypeBreakObject)
}

func (this *BreakObject) Opposite() (Object, error) {
	return nil, fmt.Errorf("BreakObject.Opposite -> unsupported")
}

func (this *BreakObject) Not() (Object, error) {
	return nil, fmt.Errorf("BreakObject.Opposite -> unsupported")
}

func (this *BreakObject) Calc(op *token.Token, right Object) (Object, error) {
	return nil, fmt.Errorf("BreakObject.Opposite -> unsupported")
}

func (this *BreakObject) Call(args []Object, insideLoop bool) (Object, error) {
	return nil, fmt.Errorf("BreakObject.Call -> unsupported")
}

func (this *BreakObject) True() bool {
	return false
}

func (this *BreakObject) Return() (bool, Object) {
	return false, nil
}

func (this *BreakObject) Break() (bool, int) {
	this.count++
	return true, this.count
}

func (this *BreakObject) calcInteger(op *token.Token, left *Integer) (Object, error) {
	return nil, fmt.Errorf("BreakObject.calcInteger -> unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *BreakObject) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	return nil, fmt.Errorf("BreakObject.calcBoolean -> unsupported op %v(%v)", op.Literal, op.Type)
}

func (this *BreakObject) calcNull(op *token.Token, left *Null) (Object, error) {
	return nil, fmt.Errorf("BreakObject.calcNull -> unsupported op %v(%v)", op.Literal, op.Type)
}
