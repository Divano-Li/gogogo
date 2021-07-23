package main

import "fmt"

func mergeSort2(a []int, head, last int) {
	if head < last {
		mid := (head + last) / 2
		mergeSort2(a, head, mid)
		mergeSort2(a, mid+1, last)
		merges(a, head, mid, last)
	}
}

func merges(a []int, head, mid, last int) {
	temp := []int{}
	i, j := head, mid+1
	for i <= mid && j <= last {
		if a[i] < a[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, a[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		temp = append(temp, a[i])
	}
	for ; j <= last; j++ {
		temp = append(temp, a[j])
	}
	for _, v := range temp {
		a[head] = v
		head++
	}
}

func main() {
	a := []int{1, 3, 6, 5, 2, 4, 10, 8, 9, 7}
	mergeSort2(a, 0, len(a)-1)
	fmt.Println(a)
}
