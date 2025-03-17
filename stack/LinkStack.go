package stack

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

type ILink interface {
	Size() int
	Pop() *Node
	Push(node *Node)
}

type LinkStack struct {
	length int
	head   *Node
}

func (stack *LinkStack) Size() int {
	return stack.length
}

func (stack *LinkStack) Pop() *Node {
	old := stack.head
	stack.head = old.Next
	stack.length--
	return old
}

func (stack *LinkStack) Push(node *Node) {
	old := stack.head
	stack.head = node
	node.Next = old
	stack.length++
}

func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func ValidLinkStack() {
	s := []int{0, 1, 2, 3, 4, 5}
	q := NewLinkStack()
	for _, val := range s {
		node := &Node{val, nil}
		q.Push(node)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for q.Size() > 0 {
		fmt.Println(q.Pop())
		fmt.Printf("--------q：%v--------\n", q)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for i := len(s) - 1; i >= 0; i-- {
		node := &Node{s[i], nil}
		q.Push(node)
	}
	fmt.Printf("--------q：%v--------\n", q)
	for q.Size() > 0 {
		fmt.Println(q.Pop())
		fmt.Printf("--------q：%v--------\n", q)
	}
}

func (stack *LinkStack) String() string {
	result := ""
	for cur := stack.head; cur != nil; cur = cur.Next {
		result += fmt.Sprintf("%v->", cur.Val)
	}
	return result
}

func (stack *LinkStack) Delete(index int) {
	if index >= stack.length {
		return
	}
	if index == 0 {
		stack.head = stack.head.Next
		return
	}
	i := 0
	for cur := stack.head; cur != nil; cur = cur.Next {
		if i == index-1 {
			cur.Next = cur.Next.Next
			break
		}
		i++
	}
}

func (stack *LinkStack) Find(val interface{}) bool {
	for cur := stack.head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			return true
		}
	}
	return false
}
