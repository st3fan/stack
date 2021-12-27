package stack_test

import (
	"testing"

	"github.com/st3fan/stack"
)

func Test_SynchronizedStackImplementsStack(t *testing.T) {
	var _ stack.StackInterface[int] = &stack.SynchronizedStack[int]{}
}
