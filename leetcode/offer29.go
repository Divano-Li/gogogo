package main

import "fmt"

func main(){
	/*a :=make([][]int,3)
	for  i,_ := range a {
		a[i] = make([]int, 3)
	}
	fmt.Print(len(a))
	fmt.Print(len(a[0]))*/
	a := []int{1,2,3}
	b := a[1:]
	c := a[0:0]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(c))
}

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	res := make([]int, 0)
	k := len(matrix)-1
	m := len(matrix[0])-1
	for i,j:=0,0;i <= k &&j <=m; i,j,k,m = i+1, j+1, k-1,m-1{
		for a,b:= i,j; b <=m ; b++ {
			res = append(res, matrix[a][b])
		}
		for a,b:= i+1, m ; a <= k; a++ {
			res = append(res, matrix[a][b])
		}
		// 这里需要判断高度是否大于一
		if k > 0 {
			for a,b:=k,m-1; b >=j; b-- {
				res = append(res, matrix[a][b])
			}
		}
		// 这里需要判断宽是否大于1
		if m > 0 {
			for a,b :=k-1,j ; a>i; a-- {
				res = append(res, matrix[a][b])
			}
		}

	}
	return res
}