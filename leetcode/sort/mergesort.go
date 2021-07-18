package main

import "fmt"

//归并排序
func mergesort(arr []int, head int, last int) {
	if head < last {
		mid := (head + last) / 2
		mergesort(arr, head, mid)
		mergesort(arr, mid+1, last)
		merge(arr, head, mid, last)
	}
}

func merge(arr []int, head int, mid int, last int) {
	temp := []int{}
	i, j := head, mid+1
	for i <= mid && j <= last {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i <= mid { //可以优化成 temp = append(temp, arr[i: mid+1])
		temp = append(temp, arr[i])
		i++
	}
	for j <= last { //可以优化成 temp = append(temp, arr[j: last+1])
		temp = append(temp, arr[j])
		j++
	}
	for _, v := range temp {
		arr[head] = v
		head++
	}
}

func main() {
	a := []int{-1, -3}
	mergesort(a, 0, len(a)-1)
	fmt.Println(a)
}
