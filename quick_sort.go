package main

import "fmt"

func quickSort(arr []int, head int, last int) {
    if head < last {
		a := arr[head]
		l,r := head, last
		for ;l < r; {
			for ;l<r && arr[r] >= a; {
				r--
			}
			if l<r {
				arr[l] = arr[r]
				l++
			}
			for ; l<r && arr[l] <= a; {
				l++
			}
			if l < r {
				arr[r] = arr[l]
				r--
			}
		}
		arr[l] = a
		quickSort(arr, head, l-1)
		quickSort(arr, l+1, last)
	}
}

func main() {
	a := []int{5,4,3,2,1}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
