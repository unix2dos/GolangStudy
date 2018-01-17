#### question
* go 版本更新 

brew update
brew install go

* cannot find package "fmt" in any of:
unset GOROOT

* flag provided but not defined: -goversion

一个是版本原因, 一个是vscode也要修改配置gopath, 坑爹

Thank you, I was able to solve this by running brew uninstall --force go and then downloading the latest installer. Anyone who reads this and wants to use brew you could probably just do brew install go after the forced uninstall. I had to restart my terminal and Gogland after doing this.

* vscode not jump define

```
        "go.useLanguageServer": true,
        "go.docsTool": "gogetdoc",
```


#### sql

1. redis
github.com/garyburd/redigo/redis

key *  //查看redis所有key

del key //删除key

type key // 获取key type

操作:
string | get key, set key
list | lrange key 0 10, lpush key value
set | smembers key, sadd key value
zset | zrange key 0 10 withscores, zadd key value
hash | hgetall key | hset key value


2. xorm
https://github.com/go-xorm/xorm



### data
set 可以用 gopkg.in/fatih/set.v0


### method

1. 方法转行成普通函数调用

```
type File struct {
	fd int
}

func OpenFile(name string) (f *File, err error) {
	// ...
}
func (f *File) Close() error {
	// ...
}
func (f *File) Read(int64 offset, data []byte) int {
	// ...
}
```

```
//转换成普通函数调用
var CloseFile = (*File).Close
var ReadFile = (*File).Read
f, _ := OpenFile("foo.dat")
ReadFile(f, 0, data)
CloseFile(f)
```

``` 
//闭包简化参数
f, _ := OpenFile("foo.dat")
var Close = func Close() error {
	return (*File).Close(f)
}
var Read = func Read(int64 offset, data []byte) int {
	return (*File).Read(f, offset, data)
}
Read(0, data)
Close()
```
### inherit

1. 结构体套上匿名结构体, 如果写方法接收是匿名的,不是自己, 所以只是继承, 不能多态

```
type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()

	return p.m[key]
}
```

### interface

+ Go语言中，对于基础类型（非接口类型）不支持隐式的转换，我们无法将一个int类型的值直接赋值给int64类型的变量，也无法将int类型的值赋值给底层是int类型的新定义命名类型的变量。Go语言对基础类型的类型一致性要求可谓是非常的严格，但是Go语言对于接口类型的转换则非常的灵活。对象和接口之间的转换、接口和接口之间的转换都可能是隐式的转换。
 
```
var (
	a io.ReadCloser = (*os.File)(f) // 隐式转换, *os.File 类型满足了 io.ReadCloser 接口
	b io.Reader     = a             // 隐式转换, io.ReadCloser 满足了 io.Reader 接口
	c io.Closer     = a             // 隐式转换, io.ReadCloser 满足了 io.Closer 接口
	d io.Reader     = c.(io.Reader) // 显式转换, io.Closer 并不显式满足 io.Reader 接口
)
```

+ 由于转换太灵活, 需要限制 >>>在protobuf中，Message接口也采用了类似的方法，也定义了一个特有的ProtoMessage，用于避免其它类型无意中适配了该接口：

```
type proto.Message interface {
	Reset()
	String() string
	ProtoMessage()
}
```
+ 通过嵌入匿名接口或嵌入匿名指针对象来实现继承的做法其实是一种纯虚继承，我们继承的只是接口指定的规范，真正的实现在运行的时候才被注入。

```
//通过将对象隐式转换为testing.TB接口类型（因为内嵌了匿名的testing.TB对象，因此是满足testing.TB接口的
type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
```

### Goroutine

* 用锁 sync.Mutex

 ```
 	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
 ```
* sync/atomic

```
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
```
* 实现 sync.once

```
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

* sync.once 实现单例模式

```
var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

* atomic.Value原子对象提供了Load和Store两个原子方法，分别用于加载和保存数据，返回值和参数都是interface{}类型，因此可以用于任意的自定义复杂类型。


### 顺序一致性内存模型

```
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {}  //有可能死循环, 也有可能空字符串
	print(a)
}
```

* 在Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性内存模型，需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两个事件是并发的。
* 解决方案:

```
func main() {
	done := make(chan int)

	go func(){
		println("你好, 世界")
		done <- 1
	}()

	<-done
}
```

* 要注意的是，在main.main函数执行之前所有代码都运行在同一个goroutine中，也是运行在程序的主系统线程中。如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到。

因为所有的init函数和main函数都是在主线程完成，它们也是满足顺序一致性模型的。

* 无缓存的Channel上的发送操作总在对应的接收操作完成前发生.
* 从无缓冲信道进行的接收，发生在对该信道进行的发送完成之前。


