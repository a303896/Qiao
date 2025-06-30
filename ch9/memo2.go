package ch9

import "sync"

// 定义生成缓存方法
type Func2 func(x string) (interface{}, error)

// 缓存结构
type Memo2 struct {
	f     Func2
	mu    sync.Mutex
	cache map[string]*entry
}

type entry struct {
	res   Result
	ready chan struct{}
}

func NewMemo2(f Func2) *Memo2 {
	return &Memo2{f: f, cache: make(map[string]*entry)}
}

// 获取缓存
func (memo *Memo2) Get(key string) Result {
	memo.mu.Lock()
	cache, ok := memo.cache[key]
	if !ok {
		cache = &entry{ready: make(chan struct{})}
		memo.cache[key] = cache
		memo.mu.Unlock()
		cache.res.value, cache.res.err = memo.f(key)
		close(cache.ready) //获取结果后，关闭通道
	} else {
		memo.mu.Unlock()
		<-cache.ready //阻滞线程，等待结果
	}
	return cache.res
}
