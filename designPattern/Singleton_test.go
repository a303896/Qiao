package designPattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestWork(t *testing.T) {
	GetInstance().Work()
}

func TestGetInstance(t *testing.T) {
	w := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			fmt.Printf("%d:%p\n", i, GetInstance())
			w.Done()
		}()
	}
	w.Wait()
}

func TestGetInstance1(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("%d:%p\n", i, GetInstanceMulti1())
		}()
	}
}

func TestGetInstance2(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("%d:%p\n", i, GetInstanceMulti2())
		}()
	}
}

func TestGetInstance3(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("%d:%p\n", i, GetInstanceMulti3())
		}()
	}
}
