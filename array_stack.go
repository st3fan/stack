package stack

type ArrayStack[T any] struct {
	data []T
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{}
}

func NewArrayStackOf[T any](elements ...T) *ArrayStack[T] {
	return &ArrayStack[T]{
		data: elements,
	}
}

func (s *ArrayStack[T]) Push(e T) bool {
	s.data = append(s.data, e)
	return true
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var empty T
		return empty, false
	}
	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e, true
}

func (s *ArrayStack[T]) Peek() (e T, found bool) {
	if len(s.data) == 0 {
		var empty T
		return empty, false
	}
	return s.data[len(s.data)-1], true
}

func (s *ArrayStack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack[T]) Clear() {
	s.data = nil
}
