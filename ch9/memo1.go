package ch9

// 定义生成缓存方法
type Func func(x string) (interface{}, error)

// 缓存结构
type Memo1 struct {
	f     Func
	cache map[string]Result
}

func New(f Func) *Memo1 {
	return &Memo1{f: f, cache: make(map[string]Result)}
}

// 获取缓存
func (memo *Memo1) Get(key string) Result {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res
}
