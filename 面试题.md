#### defer

```
func fun(i int) {
	fmt.Println(i)
}
func main() {
	a := 1
	defer fun(a) //此处输出1
	a = 0
}
```

```
func main() {
	a := 1
	defer func() { fmt.Println(a) }() //输出0
	a = 0
}
```

```
func fun(i int) {
	fmt.Println(i)
}
func main() {
	a := 1
	defer func(b int) { fmt.Println(b) }(a) //输出1
	a = 0
}
```