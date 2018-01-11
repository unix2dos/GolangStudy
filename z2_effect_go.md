### Formatting
go fmt

### Commentary
每一个包都应该有一个注释

### Names

1. Package

    a. 包名简单,易懂 用小写字母

    b. 不要使用 import .

2. Getters

    a. getter   Name()

    b. setter   setName()

3. Interface

    a. 简单的加er   Reader

4. Semicolons 分号

    a. if the newline comes after a token that could end a statement, insert a semicolon

### Control structures 控制结构

a. 没有do 和 while

b. if and switch accept an optional initialization(初始化) statement

c. break and continue statements take an optional label to identify what to break or continue;
   
d. 增加了select

e. if防御编程, 没有else

f. 没有i++ i--

g. for range 可以遍历 utf-8

h. switch加强 和if---elseif---elseif----else, 所以默认不穿透

i. break 可以  break Label

j. type switch  |   var t interface{}   t = func()  switch t:= t.(type) {}

### Functions
    
a. 可以多返回值

b. 返回值可以加参数

c. defer 功能  后进先出


### Data

1. new  return type *T  it returns a pointer to a newly allocated zero value of type T.

2. func() { return &{} }    虽然本地但会自动new

3. The expressions new(File) and &File{} are equivalent.

4.  The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).

5. Remember that make applies only to maps, slices and channels and does not return a pointer. To obtain an explicit pointer allocate with new or take the address of a variable explicitly.

6. Array 是value 

7. Map 可以用一切作为key, 除了Slice, 因为Slice没有相等比较

8. Print..   print 和 println 不需要格式化   只有f format

    %+v  key也打出    %#v  go表现打出

    ```
        fmt.Printf("Hello %d\n", 23)
        fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
        fmt.Println("Hello", 23)
        fmt.Println(fmt.Sprint("Hello ", 23))
    ```

### Initialization

1. iota

```
type ByteSize float64
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```


### Methods

1. 接收器是指针还是值

The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.


### Interfaces and other types

1. 通过类型转换我们就不用需要实现interface, 直接共享其他类型已经实现的

```
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

2. Type assertions 类型断言

```
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```

If the type assertion(断言) fails, str will still exist and be of type string, but it will have the zero value, an empty string.

3. 除了指针和接口以后外都能定义方法, 包括函数

```
type HandlerFunc func(ResponseWriter, *Request)
// ServeHTTP calls f(c, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
	f(w, req)
}
```


### The blank identifie 空白符

1. 多重赋值

2. 使用副作用

3. 检查接口


### Embedding 植入

1. 接口内嵌

// ReadWriter 接口结合了 Reader 和 Writer 接口。
type ReadWriter interface {
	Reader
	Writer
}

2. Struct内嵌

```
type ReadWriter struct {
	*Reader  // *bufio.Reader
	*Writer  // *bufio.Writer
}
```

a. 不要提供字段名, 提供了就不转发了, 除非定义再转发, 这种叫做内嵌, 不是子类

b. 还有种区分内嵌与子类的重要手段。当内嵌一个类型时，该类型的方法会成为外部类型的方法， 但当它们被调用时，该方法的接收者是内部类型，而非外部的。在我们的例子中，当 bufio.ReadWriter 的 Read 方法被调用时， 它与之前写的转发方法具有同样的效果；接收者是 ReadWriter 的 reader 字段，而非 ReadWriter 本身。



### Concurrency 并发