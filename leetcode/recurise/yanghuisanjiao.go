package main

import "fmt"

//f(i,j)=f(i−1,j−1)+f(i−1,j)
//f(i,j)=1wherej=1orj=i

func generate(numRows int) [][]int {
	a := [][]int{}
	for i:=0; i< numRows ; i++ {
		b:= []int{}
		for j:=0 ; j< i+1; j++ {
			b = append(b,getValue(i,j))
		}
		a = append(a, b)
	}
	return a
}

func getValue(i int, j int) int {
	if i== j || j == 0{
		return 1
	}
	return getValue(i-1, j-1) +  getValue(i-1, j)
}

func generate1(numRows int) [][]int {
	if numRows ==  1  {
		return [][]int{[]int{1}}
	}
	if numRows ==0 {
		return nil
	}
	a := [][]int{[]int{1}}
	for i:=1; i< numRows ; i++ {
		b:= []int{}
		for j:=0 ; j<= i; j++ {
			if i==j || j == 0 {
				b = append(b,1)
			} else {
				b = append(b, a[i-1][j-1] + a[i-1][j])
			}
		}
		a = append(a, b)
	}
	return a
}

func main () {
	b := [][]int{[]int{1}}
	fmt.Println(b[0:0])
}