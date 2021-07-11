package main

import (
	"fmt"
)

var c = make(chan int, 1)

func f() {
	c <- 'c'
	fmt.Println("在goroutine 内部")
}

func main() {
	go f()
	// time.Sleep(1 * time.Second)  加上了就会deadlock 不加上主协程先于子协程执行
	c <- 'c'
	<- c
	fmt.Println(1)
	<- c
	fmt.Println("在外部调用")

	 a:= []int{1,2,3,4}
	 fmt.Println(a[4:])
	 a = append(a[0:3], a[4:]...)
}


type A struct {
	name string
	age  int
}