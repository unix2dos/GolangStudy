package main

func fib(c, quit chan int) {
	x, y := 1, 1
	for { //这里必须加for
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() { //这里必须加go,要不然阻塞
		for i := 0; i < 10; i++ {
			println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit) //这里不能加go,一加不阻塞, 程序直接退出了
}
