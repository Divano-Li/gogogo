package main

import "fmt"

func writeRoutine(test_chan chan int, value int) {

	test_chan <- value
}

func readRoutine(test_chan chan int) {

	<-test_chan

	return
}

// 任何在主协程中先进行无缓冲chan的读写的操作都会阻塞主协程
// 为了不阻塞主协程， 需要在主协程前开辟新的协程，监听chan的读写
// 也可以利用这一特性， 在主协程中监听chan, 等到所有的子协程都执行完毕后向该chan发送给数据，
// 主协程进行下一步操作或者关闭， 这样关闭主协程可以不用sleep 的方式关闭，而是通过chan来进行通信
func main() {

	c := make(chan int)

	x := 100

	//readRoutine(c)
	//go writeRoutine(c, x)

	//writeRoutine(c, x)
	//go readRoutine(c)

	// 开辟了一个协程， 这个协程自己被阻塞了，但是下面的主协程可以继续执行，
	// 所以一旦下面的主协程向chan写入了数据， 上面的读协程完成读取操作
	go readRoutine(c)
	writeRoutine(c, x)

	// 同样的， 先开辟一个写协程， 写协程被阻塞了，等下面的读主协程开始读了，那么写协程就不会阻塞了
	go writeRoutine(c, x)
	readRoutine(c)

	fmt.Println(x)
}
