package skiplist

import "math/rand/v2"

type SkipList struct {
	head *Node
}

type Node struct {
	next     []*Node
	key, val int
}

func (s *SkipList) Search(key int) (*Node, bool) {
	current := s.head
	length := len(s.head.next)
	for i := length - 1; i >= 0; i-- { //向下遍历
		if current.key == key {
			return current, true
		}
		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
			current = current.next[i]
		}
	}
	return nil, false
}

func (s *SkipList) Insert(key int, val int) {
	current := s.head
	length := len(s.head.next)
	for i := length - 1; i >= 0; i-- { //向下遍历
		if current.key == key {
			current.val = val //key已存在更新节点val
			return
		}
		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
			current = current.next[i]
		}
	}
	level := roll()
	newNode := &Node{make([]*Node, 0), key, val}
	//新节点层级超出跳表最大高度
	if level >= length {
		s.head.next = append(s.head.next, newNode)
	}

}

//获取节点层级
func roll() int {
	level := 0
	for rand.Int()%2 > 0 {
		level++
	}
	return level
}
