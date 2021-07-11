package main

import (
	"fmt"
	"math"
	"strconv"
)

type Queue struct {
	arr []string
	head int
	tail int
}

func ConstructQueue(k int) Queue {
	a:= Queue{make([]string, k), -1, -1}
	return a
}

func (this *Queue) isEmpty() bool {
	if this.head == -1 && this.tail == -1 {
		return true
	}
	return false
}

func (this *Queue) isFull() bool {
	if this.tail == len(this.arr) - 1 {
		return true
	}
	return false
}

func (this *Queue) enQueue(value string) bool {
	if this.isFull() {
		return false
	}
	if this.isEmpty() {
		this.tail = 0
		this.head = 0
		this.arr[0] = value
		return true
	} else {
		this.tail += 1
		this.arr[this.tail] = value
		return true
	}
}

func (this *Queue) deQueue() bool {
	if this.isEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		this.head = this.head + 1
	}
	return true
}


func (this *Queue) front() string {
	if this.isEmpty() {
		return ""
	}
	return this.arr[this.head]
}


func (this *Queue) rear() string {
	if this.isEmpty() {
		return ""
	}
	return this.arr[this.tail]
}

func numSquares(n int) int {
	a:= math.Floor(math.Sqrt(float64(n)))
	if a*a == float64(n) {
		return 1
	}
	queue := ConstructQueue(100000)
	count :=0
	key:= 0
	queue.enQueue(strconv.Itoa(0))
	for queue.isEmpty() == false {
		count++
		len:= queue.tail - queue.head +1
		for i:= 0; i< len; i++ {
			key,_ = strconv.Atoi(queue.front())
			queue.deQueue()
			for j:=1; j<= int(a); j++ {
				if key + j*j == n {
					return count
				} else if key + j*j < n {
					queue.enQueue(strconv.Itoa(key + j*j ))
				}
			}
		}
	}
	return count
}

func main() {
	a := numSquares(7168)
	fmt.Println(a)
	b := []int{1,2,3}
	c:=b[0:0]
	fmt.Println(c)
}