package main

var ch chan int


func main()  {
   // var ch1 chan int
	ch = make(chan int)
	flag := "one"
	go printNum(flag)
	//go printNum("two")
	//ch1  <- 1   // 只有这一个 goroutine 的情况下， 这个channel 将会被阻塞， 所以需要从另一个goroutine  中读， 不然这个程序一直阻塞在这里
	//ch2 := make(chan int)
	/*a := <- ch
	println(a)
	b := <- ch
	println(b)*/
	L:for{
		    a := <- ch
		    println(a)
		    if a ==9 {
		    	break L
			}

		}
	}


func printNum(flag string) {
	/*for i:=0;i<100;i++ {
		fmt.Println(flag + strconv.Itoa(i))
	}*/
	ch <- 1
	for i:=0;i<10;i++ {
		ch <- i
	}
}