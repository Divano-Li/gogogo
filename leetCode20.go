package main

import (
	"fmt"
)

type StringStack struct {
	arr []string
}

func Constructor1() StringStack {
	a := StringStack{[]string{}}
	return a
}

func (this *StringStack) push(s string) {
	this.arr = append(this.arr, s)
}

func (this *StringStack) pop() {
	if len(this.arr) != 0 {
		this.arr = this.arr[0: len(this.arr)-1]
	}
}

func (this *StringStack) top() string {
	if len(this.arr) == 0 {
		return ""
	} else {
		return this.arr[len(this.arr)-1]
	}
}

func isValid(s string) bool {
  if len(s) == 0 {
  	return false
  }
  stack := Constructor1()
  for _,v:= range s {
  	if string(v) == "(" || string(v) == "{" || string(v) ==  "[" {
  		stack.push(string(v))
	} else {
		if stack.top() == "(" &&  string(v) == ")"{
			stack.pop()
			continue
		} else if stack.top() == "{" &&  string(v) == "}" {
			stack.pop()
			continue
		}  else if stack.top() == "[" &&  string(v) == "]" {
			stack.pop()
			continue
		} else {
			return false
		}
	}
  }
  if stack.top() != "" {
  	return false
  } else {
  	return true
  }

}


func main() {
	fmt.Println(isValid("({[]})"))
	/*fmt.Println(strconv.Itoa(40))
	fmt.Println(string(40))*/
}


