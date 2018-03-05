package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Liuwei struct {
	Person
	age int
}

type XuanYuan struct {
	Person
	sex string
}

func (p Person) getName() {
	fmt.Println(p.name)
}

func main() {

	l := Liuwei{Person{"liuwei"}, 10}
	x := XuanYuan{Person{"xuanyuan1"}, "nv"}

	l.getName()
	x.getName()
}
