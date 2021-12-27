package stack_test

import (
	"testing"

	"github.com/st3fan/stack"
)

func Test_ArrayStackImplementsStack(t *testing.T) {
	var _ stack.StackInterface[int] = &stack.ArrayStack[int]{}
}

func Test_NewArrayStack(t *testing.T) {
	if s := stack.NewArrayStack[int](); s == nil {
		t.Fail()
	}
}

func Test_ArrayStackOf(t *testing.T) {
	if s := stack.NewArrayStackOf(1, 2, 3); s == nil {
		t.Fail()
	}

	if s := stack.NewArrayStackOf("a", "b", "c"); s == nil {
		t.Fail()
	}

	type Thing struct{}

	if s := stack.NewArrayStackOf(&Thing{}, &Thing{}, &Thing{}); s == nil {
		t.Fail()
	}

	if s := stack.NewArrayStackOf(Thing{}, Thing{}, Thing{}); s == nil {
		t.Fail()
	}
}

func Test_PushPop1(t *testing.T) {
	s := stack.NewArrayStack[int]()
	if s == nil {
		t.Fail()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if e, found := s.Pop(); !found || e != 3 {
		t.Fail()
	}

	if e, found := s.Pop(); !found || e != 2 {
		t.Fail()
	}

	if e, found := s.Pop(); !found || e != 1 {
		t.Fail()
	}

	if _, found := s.Pop(); found {
		t.Fail()
	}
}

func Test_PushPop2(t *testing.T) {
	s := stack.NewArrayStackOf(1, 2, 3)
	if s == nil {
		t.Fail()
	}

	if e, found := s.Pop(); !found || e != 3 {
		t.Fail()
	}

	if e, found := s.Pop(); !found || e != 2 {
		t.Fail()
	}

	if e, found := s.Pop(); !found || e != 1 {
		t.Fail()
	}

	if _, found := s.Pop(); found {
		t.Fail()
	}
}

func Test_Empty1(t *testing.T) {
	s := stack.NewArrayStack[int]()
	if s == nil {
		t.Fail()
	}

	if !s.Empty() {
		t.Fail()
	}

	s.Push(42)

	if s.Empty() {
		t.Fail()
	}
}

func Test_Empty2(t *testing.T) {
	s := stack.NewArrayStackOf(1)
	if s == nil {
		t.Fail()
	}

	if s.Empty() {
		t.Fail()
	}

	if _, found := s.Pop(); !found {
		t.Fail()
	}

	if !s.Empty() {
		t.Fail()
	}
}

func Test_Clear(t *testing.T) {
	s := stack.NewArrayStackOf(1, 2, 3)

	s.Clear()

	if !s.Empty() {
		t.Fail()
	}

	if _, found := s.Pop(); found {
		t.Fail()
	}
}
