package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for {
			select {
				case data1 := <- chan1:
					fmt.Println("从chan1读取数据成功，data1:" , data1)
				case data2 := <- chan2:
					fmt.Println("---------------从chan2读取数据成功， data2:", data2)
				default:
					fmt.Println("defallt 分支called")
					time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		for i:=0; i<3; i++ {
			chan1 <- i
		}
	}()

	go func() {
		for i:=0; i<3; i++ {
			chan2 <- i
		}
	}()

	fmt.Println("liln ")
	time.Sleep(10 * time.Second)

}
