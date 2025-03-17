package bank

import (
	"fmt"
	"testing"
)

func TestInit(tet *testing.T) {
	fmt.Println(pc)
}

func TestOnce(tet *testing.T) {
	ch := make(chan int)
	//go循环变量快照问题 已在1.22修复
	//从之前的共享变量，调整为每次执行循环步都会重新实例化变量，这样避免了闭包中共享变量导致的并发问题
	//以下代码与注释部分代码结果均为正确
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("第%d轮获取的值是:%d\n", i+1, getPc(i))
			ch <- 1
		}()
		//go func(i int) {
		//	fmt.Printf("第%d轮获取的值是:%d\n", i+1, getPc(i))
		//	ch <- 1
		//}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
