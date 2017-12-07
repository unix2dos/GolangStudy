package main

import (
	"fmt"
)

type Men interface {
	say()
	walk()
}

type A struct {
}
type B struct {
}

func (a A) say() {
	fmt.Println("a say")
}
func (a A) walk() {
	fmt.Println("a walk")
}
func (a B) say() {
	fmt.Println("b say")
}
func (a B) walk() {
	fmt.Println("b walk")
}

func main() {
	arr := make([]Men, 3)
	arr[0] = A{}
	arr[1] = A{}
	arr[2] = B{}

	for _, v := range arr {
		v.say()
		v.walk()
	}
}
