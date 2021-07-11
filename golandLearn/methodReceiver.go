package main

import "fmt"

type Integer int

type Person struct {
	name string
	age int
}

func main() {
	var a Integer = 1
	var b Integer = 2
	//var i interface{} = a
	//sum := i.(Integer).Add(b)
	a.Add(b)
	fmt.Println(1)
	var p Person = Person{"lilin", 28}
	p.Older()
	fmt.Println(p)
	(&p).Older()
	fmt.Println(p)
	p.Older1()
	fmt.Println(p)
	(&p).Older1()
	fmt.Println(p)
}


func (a *Integer) Add(b Integer) Integer {

	return *a + b
}

func (p *Person) Older() {
	p.age = p.age + 1
}

func (p Person) Older1() {
	p.age = p.age + 1
}

