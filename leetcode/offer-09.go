package main


type CQueue struct {
  in stack
  out stack
}


/*func Constructor() CQueue {
   return CQueue{}
}*/


func (this *CQueue) AppendTail(value int)  {
	this.in.push(value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.out) != 0 {
    	return this.out.pop()
	} else if len(this.in) !=0 {
		for len(this.in) !=0 {
			this.out.push(this.in.pop())
		}
		return this.out.pop()
	}
	return -1
}

type stack []int

func (this *stack) push(value int) {
	*this = append(*this,value)
}

func (this *stack) pop() int {
	len := len(*this)
	if len == 0 {
		return -1
	}
	res := (*this)[len-1]
	*this = (*this)[:len-1]
	return res
}




/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */