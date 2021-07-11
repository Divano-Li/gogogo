package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("hello world")
	fmt.Println(1/10)
	a := buildStudent()
	fmt.Printf("%v\n", a)

}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

type Student struct {
	name string
	age  int
}

func buildStudent() *Student {
	a := new(Student)
	a.name = "tianping"
	a.age = 200
	return a
}

func (h *Student) getName() string {
	return h.name
}
