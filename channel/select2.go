//超时处理
package main

import (
	"time"
)

//要多次练习接收发送
//接收 <- channel
//发送 channel <- 数据
//接收 <- channel
//发送 channel <- 数据
//接收 <- channel
//发送 channel <- 数据
func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() { //此处用go是为了不阻塞
		for {
			select {
			case <-c:
			case <-time.After(5 * time.Second):
				println("5秒了")
				quit <- 0
				break
			}
		}
	}()

	<-quit //此处是为了阻塞
}
