package stack_test

import (
	"lsroa/go-machine/pkg/stack"
	"testing"
)

func expectEq[T comparable](t *testing.T, a, b T) {
	t.Helper()

	if a != b {
		t.Errorf("\ngot:  %v \nwant: %v", a, b)
	}

}

func TestStack(t *testing.T) {
	list := stack.New[int]()
	list.Push(5)
	list.Push(7)
	list.Push(9)

	expectEq(t, list.Pop(), 9)
	expectEq(t, list.Count, 2)

	list.Push(11)
	expectEq(t, list.Pop(), 11)
	expectEq(t, list.Pop(), 7)
	expectEq(t, list.Pop(), 5)

	expectEq(t, list.Pop(), 0)
}
