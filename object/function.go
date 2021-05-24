package object

import (
	"Q/function"
	"Q/token"
	"fmt"
)

// Function : implement Object
type Function struct {
	Fn  function.Function
	Env *Env
}

func (this *Function) Type() ObjectType {
	return ObjectTypeFunction
}

func (this *Function) Inspect() string {
	return this.Fn.Inspect()
}

func (this *Function) Not() (Object, error) {
	return nil, fmt.Errorf("Function.Not not supported")
}

func (this *Function) Opposite() (Object, error) {
	return nil, fmt.Errorf("Function.Opposite not supported")
}

func (this *Function) Calc(op *token.Token, right Object) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.Calc not supported")
}

func (this *Function) True() bool {
	return false
}

func (this *Function) Return() (bool, Object) {
	// TODO
	return false, nil
}

func (this *Function) calcInteger(op *token.Token, left *Integer) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.calcInteger not supported")
}

func (this *Function) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.calcBoolean not supported")
}

func (this *Function) calcNull(op *token.Token, left *Null) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.calcNull not supported")
}
