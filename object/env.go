package object

import "fmt"

type Env struct {
	outer *Env
	m     map[string]Object
}

func NewEnv() *Env {
	return &Env{m: map[string]Object{}}
}

func (this *Env) Get(name string) (Object, bool) {
	v, ok := this.m[name]
	if !ok && nil != this.outer {
		v, ok = this.outer.Get(name)
		return v, ok
	}
	return v, ok
}

func (this *Env) Set(name string, val Object) Object {
	this.m[name] = val
	return val
}

func (this *Env) Assign(name string, val Object) error {
	if _, ok := this.m[name]; ok {
		this.m[name] = val
	} else {
		return fmt.Errorf("Env.Assign -> `%v` undefined", name)
		// TODO
		// if nil == this.outer {
		// 	return fmt.Errorf("Env.Assign -> `%v` undefined", name)
		// } else {
		// 	return this.outer.Assign(name, val)
		// }
	}
	return nil
}

func newEnclosedEnv(outer *Env) *Env {
	env := NewEnv()
	env.outer = outer
	return env
}

func newFunctionEnv(outer *Env, args []string, values []Object) *Env {
	env := newEnclosedEnv(outer)
	for i, name := range args {
		env.Set(name, values[i])
	}
	return env
}
