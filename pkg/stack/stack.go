package stack

type Stack[T any] struct {
	Count int
	head  *Node[T]
}

func New[T any]() *Stack[T] {
	node := &Node[T]{}
	return &Stack[T]{Count: 0, head: node}
}

func (s *Stack[T]) Push(item T) {
	s.Count = s.Count + 1

	// Create node pointing to list's head
	node := &Node[T]{
		Data: item,
		Next: s.head,
	}
	// Assign new node as list's head
	s.head = node
}

func (s *Stack[T]) Pop() T {
	s.Count = s.Count - 1

	node := s.head
	s.head = node.Next

	node.Next = nil
	return node.Data
}

type Node[T any] struct {
	Data T
	Next *Node[T]
}
