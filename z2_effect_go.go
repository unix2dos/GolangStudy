package main

import "fmt"

type A struct {
}

func (a A) a() {
	fmt.Println("aa")
}

type B struct {
}

func (b B) b() {
	fmt.Println("bb")
}

type ALL struct {
	a *A
	*B
}

func main() {
	all := &ALL{new(A), new(B)}
	// all.a()
	all.b()
}
