package object

type Env struct {
	m map[string]Object
}

func NewEnv() *Env {
	return &Env{m: map[string]Object{}}
}

func (this *Env) Get(name string) (Object, bool) {
	v, ok := this.m[name]
	return v, ok
}

func (this *Env) Set(name string, val Object) Object {
	this.m[name] = val
	return val
}
