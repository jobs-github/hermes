package object

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
