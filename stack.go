package stack

type StackInterface[T any] interface {
	Push(e T) bool
	Pop() (T, bool)
	Peek() (T, bool)
	Empty() bool
	Clear()
}
