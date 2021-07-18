package main

import (
	"fmt"
)

func quick_sorts(arr []int, head int, last int) {
	if head < last {
		x := arr[head]
		i, j := head, last
		for i < j {
			for i < j && arr[j] >= x {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for i < j && arr[i] <= x {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = x
		quick_sorts(arr, head, i-1)
		quick_sorts(arr, i+1, last)
	}
}

func main() {
	a := []int{5, 2, 3, 1, 4}
	quick_sorts(a, 0, len(a)-1)
	fmt.Println(a)
}

type pool struct {
	MaxNums int
	Nums    int
	signal  chan int
}

//创建资源池
func initPool() *pool {
	return &pool{100, 0, make(chan int, 100)}
}

func loop() {
	p := initPool()
	for {
		select {
		case a := <-p.signal:
			if a == 1 {
				p.getGoroutine()
			} else {
				p.receiveGoroutine()
			}
		default:
		}
	}
}

func (p *pool) getGoroutine() bool {

	if p.Nums > p.MaxNums {
		return false
	} else {
		p.Nums = p.Nums + 1
		return true
	}
}

func (p *pool) receiveGoroutine() {
	p.Nums = p.Nums - 1
}
