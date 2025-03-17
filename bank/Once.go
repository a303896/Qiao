package bank

import (
	"fmt"
	"sync"
)

var pc [256]byte

//
//func init() {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//}

// 该部分代码可以替代3-9行
//var pc [256]byte = func() (pc [256]byte) {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//	return
//}()

// 使用once
var onceLoad sync.Once

func initPc() {
	for i := range pc {
		pc[i] = byte(i)
	}
	fmt.Println("初始化PC变量", pc)
}

func getPc(x int) byte {
	if x < 0 || x > 255 {
		return 0
	}
	onceLoad.Do(initPc)
	return pc[x]
}

type Number struct {
	Num int
}
