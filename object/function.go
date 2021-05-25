package object

import (
	"Q/function"
	"Q/token"
	"fmt"
)

// Function : implement Object
type Function struct {
	Fn       function.Function
	Args     []string
	EvalBody func(env *Env) (Object, error)
	Env      *Env
}

func (this *Function) Type() ObjectType {
	return ObjectTypeFunction
}

func (this *Function) Inspect() string {
	return this.Fn.Inspect()
}

func (this *Function) Not() (Object, error) {
	return nil, fmt.Errorf("Function.Not -> unsupported")
}

func (this *Function) Opposite() (Object, error) {
	return nil, fmt.Errorf("Function.Opposite -> unsupported")
}

func (this *Function) Calc(op *token.Token, right Object) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.Calc -> unsupported")
}

func (this *Function) Call(args []Object) (Object, error) {
	if len(args) != len(this.Args) {
		return nil, fmt.Errorf("Function.Call -> %v args provided, but %v args required", len(args), len(this.Args))
	}
	innerEnv := newFunctionEnv(this.Env, this.Args, args)
	evaluated, err := this.EvalBody(innerEnv)
	if nil != err {
		return nil, fmt.Errorf("Function.Call | %v", err)
	}
	if isReturn, rc := evaluated.Return(); isReturn {
		return rc, nil
	}
	return evaluated, nil
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
	return nil, fmt.Errorf("Function.calcInteger -> unsupported")
}

func (this *Function) calcBoolean(op *token.Token, left *Boolean) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.calcBoolean -> unsupported")
}

func (this *Function) calcNull(op *token.Token, left *Null) (Object, error) {
	// TODO
	return nil, fmt.Errorf("Function.calcNull -> unsupported")
}
