package queue

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

type ILink interface {
	Size() int
	Dequeue() *Node
	Enqueue(node *Node)
}

type LinkQueue struct {
	length int
	head   *Node
	tail   *Node
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{}
}

func (s *LinkQueue) Size() int {
	return s.length
}

func (s *LinkQueue) Dequeue() *Node {
	if s.length == 0 {
		return nil
	}
	old := s.head
	s.head = old.Next
	s.length--
	if s.length == 0 {
		s.tail = nil
	}
	return old
}

func (s *LinkQueue) Enqueue(node *Node) {
	if s.length == 0 {
		s.head = node
		s.tail = node
	} else {
		old := s.tail
		s.tail = node
		old.Next = node
	}
	s.length++
}

func ValidLinkQueue() {
	s := []int{0, 1, 2, 3, 4, 5}
	q := NewLinkQueue()
	for _, val := range s {
		node := &Node{val, nil}
		q.Enqueue(node)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for q.Size() > 0 {
		fmt.Println(q.Dequeue())
		fmt.Printf("--------q：%v--------\n", q)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for i := len(s) - 1; i >= 0; i-- {
		node := &Node{s[i], nil}
		q.Enqueue(node)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for q.Size() > 0 {
		fmt.Println(q.Dequeue())
		fmt.Printf("--------q：%v--------\n", q)
	}
}
