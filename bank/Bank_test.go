package bank

import (
	"fmt"
	"testing"
)

func TestExcel2string(t *testing.T) {
	ch := make(chan bool)
	go func() {
		Deposit(100)
		ch <- true
	}()
	go func() {
		Deposit(200)
		ch <- true
	}()
	select {
	case <-ch:
		fmt.Printf("当前余额为:%d", Balance())
	}
}

// 测试读锁是否共享
func TestRLock(t *testing.T) {
	go func() {
		fmt.Printf("余额:%d\n", Balance2())
	}()
	go func() {
		fmt.Printf("余额:%d\n", Balance2())
	}()
}

// 测试读写锁是否共享
// fatal error: all goroutines are asleep - deadlock!
func TestRWLock(t *testing.T) {
	fmt.Printf("余额:%d\n", Balance2())
	Deposit2(100)
	fmt.Printf("余额:%d\n", Balance2())
}
