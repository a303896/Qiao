package ch9

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var urls = []string{
	"https://www.baidu.com",
	"https://www.taobao.com",
	"https://www.jd.com",
	"https://www.baidu.com",
	"https://www.taobao.com",
	"https://www.jd.com",
}

// 定义缓存结果结构
type Result struct {
	value interface{}
	err   error
}

type Memo interface {
	Get(key string) Result
}

func httpGetBody(url string) (interface{}, error) {
	num := rand.IntN(1000)
	time.Sleep(time.Duration(num) * time.Millisecond)
	return "http请求内容" + time.Now().String(), nil
}

func getContent(memo Memo) {
	for _, url := range urls {
		start := time.Now()
		res := memo.Get(url)
		if res.err != nil {
			panic(res.err)
		}
		fmt.Printf("%s, %s, %s\n", url, res.value, time.Since(start))
	}
}

// 多线程版
func getContentMulti(memo Memo) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			res := memo.Get(url)
			if res.err != nil {
				panic(res.err)
			}
			fmt.Printf("%s, %s, %s\n", url, res.value, time.Since(start))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
