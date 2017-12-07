package main

import "fmt"

type Interface interface {
}

func main() {
	arr := make([]Interface, 4)
	arr[0] = 1
	arr[1] = 1.2
	arr[2] = true
	arr[3] = "111111111"

	for _, v := range arr {
		switch a := v.(type) {
		case int:
			fmt.Println("int ", a)
		case float64:
			fmt.Println("float64 ", a)
		case bool:
			fmt.Println("bool ", a)
		case string:
			fmt.Println("string ", a)
		}
	}
}
