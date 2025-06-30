package skiplist

import (
	"sync"
	"sync/atomic"
)

// 并发安全跳表  https://zhuanlan.zhihu.com/p/622177029
type ConcurrentSkipList struct {
	cap atomic.Int32 //跳表元素个数,通过atomic.Int32保证原子性

	DeleteMutex sync.RWMutex //一把全局的读写锁；get、put 操作取读锁，实现共享；delete 操作取写锁，实现全局互斥

	KeyMutex sync.RWMutex //每个 key 对应的一把互斥锁. 针对同一个 key 的 put 操作需要取 key 锁实现互斥

	head *NodeTS //虚拟头节点

	nodePool sync.Pool //NodeTS对象池,提高对象复用率,增加性能

	compare func(k1, k2 any) bool //比较两个NodeTS对象大小的方法
}

type NodeTS struct {
	key, val int
	next     []*NodeTS
	//每个 node 对应的一把读写锁. 在 get 检索过程中，会逐层对左边界节点加读锁；put 在插入新节点过程中，会逐层对左边界节点加写锁
	sync.RWMutex
}

func NewConcurrentSkipList(compare func(k1, k2 any) bool) *ConcurrentSkipList {
	return &ConcurrentSkipList{
		head: &NodeTS{next: make([]*NodeTS, 1)},
		nodePool: sync.Pool{
			New: func() interface{} {
				return &NodeTS{}
			},
		},
		compare: compare,
	}
}

// 查找节点的时候,需要给边界加读锁,防止在查找过程中边界节点被修改
func (s *ConcurrentSkipList) Search(key int) (*NodeTS, bool) {
	s.DeleteMutex.RLock() //首先加全局删除读锁,与删除操作互斥
	defer s.DeleteMutex.RUnlock()
	var last *NodeTS //边界节点  比如在10-20之间,查找15 10就是边界节点
	current := s.head
	length := len(s.head.next)
	for i := length - 1; i >= 0; i-- { //从顶层向下遍历
		for current.next[i] != nil && s.compare(current.next[i].key, key) { //本层级向右遍历
			current = current.next[i]
		}
		if current != last { //
			current.RLock()
			defer current.RUnlock()
			last = current
		}
		if current.next[i] != nil && current.next[i].key == key {
			return current, true
		}
	}
	return nil, false
}

// 插入节点先获取全局删除读锁,再获取边界读锁
//func (s *ConcurrentSkipList) Insert(key int, val int) {
//	if ele, ok := s.Search(key); ok {
//		ele.val = val
//		return
//	}
//	level := roll()
//	newNode := &Node{make([]*Node, level+1), key, val}
//	//新节点层级超出跳表最大高度
//	for level > len(s.head.next)-1 {
//		s.head.next = append(s.head.next, nil)
//	}
//	current := s.head
//	for i := level; i >= 0; i-- { //向下遍历
//		//本层循环结束,current为新节点插入位置
//		for current.next[i] != nil && current.next[i].key < key { //本层级向右遍历
//			current = current.next[i]
//		}
//		//新节点本层的下一个节点
//		newNode.next[i] = current.next[i]
//		current.next[i] = newNode //替换节点
//	}
//}
