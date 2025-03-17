package stack

import (
	"fmt"
	"testing"
)

func TestValidLinkStack(t *testing.T) {
	ValidLinkStack()
}

func TestDelete(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	q := NewLinkStack()
	for _, val := range s {
		node := &Node{val, nil}
		q.Push(node)
	}
	q.Delete(0)
	fmt.Println("link:", q)
	q.Delete(3)
	fmt.Println("link:", q)
}

func TestLinkStack_Find(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	q := NewLinkStack()
	for _, val := range s {
		node := &Node{val, nil}
		q.Push(node)
	}
	fmt.Println("result:", q.Find(4))
	fmt.Println("result:", q.Find(10))
}
