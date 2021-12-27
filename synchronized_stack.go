package stack

import "sync"

type SynchronizedStack[T any] struct {
	mutex sync.Mutex
	stack StackInterface[T]
}

func NewSynchronizedStack[T any](stack StackInterface[T]) *SynchronizedStack[T] {
	return &SynchronizedStack[T]{stack: stack}
}

func (s *SynchronizedStack[T]) Push(e T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stack.Push(e)
}

func (s *SynchronizedStack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stack.Pop()
}

func (s *SynchronizedStack[T]) Peek() (e T, found bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stack.Peek()
}

func (s *SynchronizedStack[T]) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stack.Empty()
}

func (s *SynchronizedStack[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack.Clear()
}
