package function

type Function struct {
	Inspect    func() string
	Arguments  func() int
	ArgumentOf func(idx int) string
	Body       func() string
}
