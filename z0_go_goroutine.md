#### 数据竞争

1. 提前初始化好变量
2. 数据在一个gorouting里修改读取(不能一个解决的串行gorouting)

```
    go func(){
        for {
            select{
            }
        }
    }()
```
3. 加锁


#### 内存同步


##### 在一个独立的goroutine中，每一个语句的执行顺序是可以被保证的，也就是说goroutine内顺序是连贯的。

但是在不使用channel且不使用mutex这样的显式同步操作时，我们就没法保证事件在不同的goroutine中看到的执行顺序是一致的了。


#### 缓存无锁
1. 为了效率加上缓存, 但是get的时候存在竞争
    
```
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	// go AA()
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func AA() {

}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func main() {
	memo := New(httpGetBody)
	go func() {
		value, err := memo.Get("https://www.baidu.com")
		fmt.Println(string(value.([]byte)), err)
	}()
	go func() {
		value, err := memo.Get("https://www.baidu.com")
		fmt.Println(string(value.([]byte)), err)
	}()

	time.Sleep(time.Second * 5)
}

```












