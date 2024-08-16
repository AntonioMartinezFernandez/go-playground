package sum

type Elements struct {
	First  int
	Second int
}

func (e *Elements) Sum() int {
	return e.First + e.Second
}
