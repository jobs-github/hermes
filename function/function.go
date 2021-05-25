package function

type Function struct {
	Inspect    func() string
	ArgumentOf func(idx int) string
	Body       func() string
}
