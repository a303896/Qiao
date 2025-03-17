package designPattern

import (
	"fmt"
	"sync"
)

var s *singleton

type singleton struct {
}

/*
*
• singleton 是包内的不可导出类型，在包外即便获取到了，也无法直接作为方法的入参或者出参进行传递，显得很呆
• singleton 的对外暴露，使得 singleton 所在 package 的代码设计看起来是自相矛盾的，混淆了 singleton 这个不可导出类型的边界和定位
*/
type Instance interface {
	Work()
}

//饿汉式单例  不管用没用到该实例,每次导入包时都会生成一个实例
//func init() {
//	if s == nil {
//		s = &singleton{}
//	}
//}

//func GetInstance() *singleton {
//	return s
//}

// 懒汉模式
func GetInstance() Instance {
	if s == nil {
		s = newInstance()
	}
	return s
}

func (*singleton) Work() {
	fmt.Println("singleton work")
}

func newInstance() *singleton {
	return &singleton{}
}

/*
*
多线程版本1
保证在多线程的情况只会产生一个实例
但是在获取实例的过程中是串行的  性能低
*/
func GetInstanceMulti1() Instance {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	if s == nil {
		s = newInstance()
	}
	return s
}

/*
*
多线程版本2
保证在多线程的情况只会产生一个实例
在实例已存在的情况能保证并发性能
只有在实例不存在的时候进行加锁操作
*/
func GetInstanceMulti2() Instance {
	if s != nil {
		return s
	}
	var m sync.Mutex
	m.Lock()
	defer m.Unlock()
	if s == nil {
		s = newInstance()
	}
	return s
}

func GetInstanceMulti3() Instance {
	m := sync.Once{}
	m.Do(func() {
		s = newInstance()
	})
	return s
}
