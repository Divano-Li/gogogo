package main

import "fmt"

func quickSort(arr []int, head int, last int) {
	if head < last {
		a := arr[head]
		l, r := head, last
		for l < r {
			for l < r && arr[r] >= a {
				r--
			}
			if l < r {
				arr[l] = arr[r]
				l++
			}
			for l < r && arr[l] <= a {
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
	a := []int{5, 4, 3, 2, 1}
	quick_sort2(a, 0, len(a)-1)
	fmt.Println(a)
	b := map[string]int{"s": 1}
	fmt.Println(b)
	fmt.Println(b["s"])
	v, ok := b["1"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(2)
	}
	c := []int{1, 2, 3, 4, 5}
	for k, v := range c {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func quick_sort2(arr []int, head int, last int) {
	if head < last {
		a := arr[head]
		l, r := head, last
		for l < r {
			for l < r && arr[r] >= a {
				r--
			}
			if l < r {
				arr[l] = arr[r]
				l++
			}
			for l < r && arr[l] <= a {
				l++
			}
			if l < r {
				arr[r] = arr[l]
				r--
			}
		}
		arr[l] = a
		quick_sort2(arr, head, l-1)
		quick_sort2(arr, l+1, last)
	}
}
