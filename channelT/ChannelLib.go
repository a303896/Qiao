package channelT

import (
	"fmt"
	"time"
)

func test() {
	defer fmt.Println("defer A")
	ch := make(chan int, 1)
	go func() {
		defer fmt.Println("defer B")
		ch <- 10 //channel内没有值会导致从channel取值时程序阻塞,发生死锁报错
	}()
	num := <-ch //此处会等待channel写操作
	fmt.Println("程序结束num:", num)
}

func Test2() {
	defer fmt.Println("主线程结束")
	ch := make(chan int, 3) //超出channel cap时会阻塞等待,直至channel被全部写入取出
	go func() {
		defer fmt.Println("子线程结束")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("写入num:", i)
		}
	}()
	time.Sleep(1 * time.Second)
	//for num := range ch {
	//	num = <-ch
	//	fmt.Println("num:", num)
	//}
	for i := 0; i < 4; i++ {
		fmt.Println("num:", <-ch)
	}
}

func Test3() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	//go func() {
	//	for i := 2; i <= 10; i++ {
	//		if i%2 == 0 {
	//			ch1 <- i
	//			ch2 <- i
	//		} else {
	//			ch3 <- i
	//		}
	//	}
	//	close(ch1)
	//	close(ch2)
	//	close(ch3)
	//}()
	//ch1 <- 2
	//ch2 <- 2
	//ch3 <- 3
	//for {
	//
	//}

	// 如果没有任意case满足条件的情况下,select会阻滞,select{}会永远地等待
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	//case条件代码会先全部执行,再选则满足条件的case块执行,若同时满足多个条件 则随机选择一个case块执行
	select {
	case <-func() chan int {
		fmt.Println("执行ch1")
		return ch1
	}():
		//fmt.Println("读取ch1:", i)
	case i := <-func() chan int {
		fmt.Println("执行ch2")
		return ch2
	}():
		fmt.Println("读取ch2:", i)
	case i := <-func() chan int {
		fmt.Println("执行ch3")
		return ch3
	}():
		fmt.Println("读取ch3:", i)
		//default:
		//	fmt.Println("default:")
	}
}
