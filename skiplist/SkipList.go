package skiplist

import (
	"fmt"
	"math/rand/v2"
)

// 跳表  https://zhuanlan.zhihu.com/p/620291031
type SkipList struct {
	head *Node //虚拟节点,不可删除
}

type Node struct {
	next     []*Node
	key, val int
}

func (s *SkipList) Search(key int) (*Node, bool) {
	current := s.head
	length := len(s.head.next)
	for i := length - 1; i >= 0; i-- { //从顶层向下遍历
		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
			current = current.next[i]
		}
		if current.next[i] != nil && current.next[i].key == key {
			return current, true
		}
	}
	return nil, false
}

func (s *SkipList) Insert(key int, val int) {
	if ele, ok := s.Search(key); ok {
		ele.val = val
		return
	}
	level := roll()
	newNode := &Node{make([]*Node, level+1), key, val}
	//新节点层级超出跳表最大高度
	for level > len(s.head.next)-1 {
		s.head.next = append(s.head.next, nil)
	}
	current := s.head
	for i := level; i >= 0; i-- { //向下遍历
		//本层循环结束,current为新节点插入位置
		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
			current = current.next[i]
		}
		//新节点本层的下一个节点
		newNode.next[i] = current.next[i]
		current.next[i] = newNode //替换节点
	}
}

// 获取节点层级
func roll() int {
	level := 0
	for rand.Int()%2 > 0 {
		level++
	}
	return level
}

func (s *SkipList) Delete(key int) {
	ele, ok := s.Search(key)
	if !ok {
		return
	}
	length := len(ele.next) //获取要删除节点的level
	current := s.head       //从该level的头节点开始遍历,因为是每层结构是单向链表,这样才能找到要删除节点的上一个节点
	for i := length - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
			current = current.next[i]
		}
		if current.next[i] == nil || current.next[i].key > key { //排除没找到要删除节点的情况(可以不用判断)
			continue
		}
		//到这就算找到要删除节点的位置,调整指针引用
		current.next[i] = current.next[i].next[i]
	}
	s.CutLevel()
}

// 修剪跳表层级
func (s *SkipList) CutLevel() {
	level := len(s.head.next) - 1
	cutNum := 0 //需要修剪的高度
	for i := level; i >= 0; i-- {
		if s.head.next[i] == nil {
			cutNum++
		} else {
			break
		}
	}
	if cutNum > 0 {
		s.head.next = s.head.next[:level-cutNum+1]
	}
}

func (s *SkipList) String() string {
	result := ""
	level := len(s.head.next) - 1
	for i := level; i >= 0; i-- {
		current := s.head.next[i]
		result += fmt.Sprintf("%d->", current.key)
		for current.next[i] != nil {
			current = current.next[i]
			result += fmt.Sprintf("%d->", current.key)
		}
		result += "\n"
	}
	return result
}

func CreateSkipList(s []int) *SkipList {
	list := &SkipList{head: &Node{}}
	for _, v := range s {
		list.Insert(v, v)
	}
	return list
}
