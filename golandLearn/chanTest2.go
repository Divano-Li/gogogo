package main

import "fmt"

func client(ch chan string, sigover chan struct{}) {
	str := "Client say hello"
	ch <- str
	str = <-ch
	fmt.Println("client recieved: " + str)
	sigover <- struct{}{}
}

func server(ch chan string) {
	str := <-ch
	fmt.Println("server recieved: " + str)
	str = "Server say world"
	ch <- str
	close(ch)
}

func main() {
	ch := make(chan string)
	sigover := make(chan struct{})
	go client(ch, sigover)
	go server(ch)
	<-sigover   // 表示收到over信号后主线程才会停， 不然主线程一直阻塞住
}
