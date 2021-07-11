package main

type MinStack struct {
	arr []int
	min []int
}

func Constructor() MinStack {
	minStack := MinStack{[]int{}, []int{}}
	return minStack
}

func (this *MinStack) Push(x int) {
	if len(this.arr) == 0 {
		this.min = append(this.min, x)
	} else {
		if this.min[len(this.min)-1] > x {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, this.min[len(this.min)-1])
		}
	}
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop() {
	if len(this.arr) != 0 {
		this.arr = this.arr[0 : len(this.arr)-1]
		this.min = this.min[0 : len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
