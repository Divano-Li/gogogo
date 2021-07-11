package main

/*
go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
在一个主进程，多个go程处理逻辑的结构中，这个很重要，如果不用recover捕获panic异常，会导致整个进程出错中断
*/
import "fmt"

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")

	fmt.Println("f")
}
