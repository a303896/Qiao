package queue

import "fmt"

type IQueue interface {
	isEmpty() bool
	enqueue(item interface{})
	dequeue() interface{}
	size() int
}

type ArrayQueue struct {
	length int
	store  []interface{}
}

func (q *ArrayQueue) isEmpty() bool {
	return q.length == 0
}

func (q *ArrayQueue) enqueue(item interface{}) {
	q.store[q.length] = item
	q.length++
}

func (q *ArrayQueue) dequeue() interface{} {
	result := q.store[0]
	copy(q.store, q.store[1:])
	q.store[q.length-1] = nil
	q.length--
	return result
}

func (q *ArrayQueue) size() int {
	return q.length
}

func ValidQueue() {
	s := []int{0, 1, 2, 3, 4, 5}
	qs := make([]interface{}, len(s))
	q := &ArrayQueue{0, qs}
	for _, val := range s {
		q.enqueue(val)
	}
	for !q.isEmpty() {
		fmt.Printf("--------q：%v--------\n", q)
		fmt.Println(q.dequeue())
	}
	fmt.Printf("--------q：%v--------\n", q)
	for i := len(s) - 1; i >= 0; i-- {
		q.enqueue(s[i])
	}
	for !q.isEmpty() {
		fmt.Println(q.dequeue())
	}
}
