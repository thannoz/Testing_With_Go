package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmtpy() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmtpy() {
		var zero T
		return zero, false
	}

	idx := len(s.values) - 1
	el := s.values[idx]
	s.values = s.values[:idx]
	return el, true
}
