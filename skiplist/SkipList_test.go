package skiplist

import (
	"fmt"
	"testing"
)

func TestCreateSkipList(t *testing.T) {
	//s := make([]int, 20)
	//for i := 0; i < 20; i++ {
	//	s[i] = rand.Intn(1000)
	//}
	s := []int{289, 191, 845, 747, 930, 455, 742, 619, 888, 427, 749, 434, 944, 99, 509, 435, 141, 376, 329, 565}
	fmt.Println(s)
	ss := CreateSkipList(s)
	fmt.Println(ss)
	ss.Delete(747)
	fmt.Println(">>>>>>>删除747节点<<<<<<")
	fmt.Println(ss)
	ss.Insert(555, 555)
	fmt.Println(">>>>>>>插入555节点<<<<<<")
	fmt.Println(ss)
	top := ss.head.next[len(ss.head.next)-1].val
	fmt.Printf(">>>>>>>删除%d节点<<<<<<\n", top)
	ss.Delete(top)
	fmt.Println(ss)
}
