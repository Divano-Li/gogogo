package queue

import (
	"strconv"
)

type Queue struct {
	arr []byte
	head int
	tail int
}

func ConstructQueue(k int) Queue {
	a:= Queue{make([]byte, k), -1, -1}
	return a
}

func (this *Queue) isEmpty() bool {
	if this.head == -1 && this.tail == -1 {
		return false
	}
	return true
}

func (this *Queue) isFull() bool {
	if this.tail == len(this.arr) - 1 {
		return true
	}
	return false
}

func (this *Queue) enQueue(value byte) bool {
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


func (this *Queue) front() byte {
	if this.isEmpty() {
		return 0
	}
	return this.arr[this.head]
}


func (this *Queue) rear() byte {
	if this.isEmpty() {
		return 0
	}
	return this.arr[this.tail]
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	a := map[string]byte{}
	ii := len(grid)
	jj := len(grid[0])
    count := 0
	key := ""
	queue := ConstructQueue(90000)
	var temp byte
	for i:=0;i< len(grid); i++ {
		for j:=0;j<len(grid[i]);j++ {
			key = strconv.Itoa(i) + strconv.Itoa(j)
			temp = grid[i][j]
			_, ok := a[key];  if !ok && temp == 1 {
				count++
				a[key] = temp
				queue.enQueue(temp)
				for ; !queue.isEmpty() ; {
					if i-1 >=0 {
						key = strconv.Itoa(i-1) + strconv.Itoa(j)
						_, ok := a[key]; if !ok && grid[i-1][j] ==1 {
							queue.enQueue(grid[i-1][j])
							a[key] = grid[i-1][j]
						}
					}
					if i+1 <= ii -1 {
						key = strconv.Itoa(i+1) + strconv.Itoa(j)
						_, ok := a[key]; if !ok && grid[i+1][j] ==1 {
							queue.enQueue(grid[i+1][j])
							a[key] = grid[i+1][j]
						}
					}
					if j-1 >= 0 {
						key = strconv.Itoa(i) + strconv.Itoa(j-1)
						_, ok := a[key]; if !ok && grid[i][j-1] ==1 {
							queue.enQueue(grid[i][j-1])
							a[key] = grid[i][j-1]
						}
					}
					if j+1 <= jj-1 {
						key = strconv.Itoa(i) + strconv.Itoa(j+1)
						_, ok := a[key]; if !ok && grid[i][j+1] ==1 {
							queue.enQueue(grid[i][j+1])
							a[key] = grid[i][j+1]
						}
					}
					queue.deQueue()
				}

			}
		}
	}
	return  count
}
