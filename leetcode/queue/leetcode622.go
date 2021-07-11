package queue

type MyCircularQueue struct {
 	arr []int
 	head int
 	tail int
}

func Constructor(k int) MyCircularQueue {
   temp := make([]int, k)
   head := -1
   tail := -1
   queue := MyCircularQueue{temp, head, tail}
   return  queue
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.tail = 0
		this.head = 0
		this.arr[0] = value
		return  true
	}
	if this.tail == len(this.arr)-1 {
		this.tail = 0
		this.arr[0] = value
	} else {
		this.tail += 1
		this.arr[this.tail] = value
	}
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	if this.head == len(this.arr) -1 {
		this.head  =  0
	} else {
		this.head += 1
	}
	return true
}


func (this *MyCircularQueue) Front() int {
	 if this.IsEmpty() {
		return -1
	 }
     return this.arr[this.head]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.tail]
}


func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == -1 && this.tail == -1  {
		return true
	} else {
		return false
	}
}


func (this *MyCircularQueue) IsFull() bool {
	if this.head == 0 && this.tail == len(this.arr) -1 {
		return true
	}
	if this.head - this.tail ==1 {
		return true
	}
	return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */