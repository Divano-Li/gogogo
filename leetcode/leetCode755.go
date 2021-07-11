package main

import "fmt"

func dailyTemperatures(T []int) []int {
	m := map[int]int{}
	result := []int{0}
	m[len(T)-1] = 0

	for i:=len(T)-2; i >=0 ;i--{
		if T[i] < T[i+1] {
			m[i] = i+1
			result = append([]int{1}, result...)
		} else {
			index := getIndex(T[i], i+1, m ,T)
			if index == 0 {
				m[i] = 0
				result = append([]int{0}, result...)
			} else {
				m[i] = index
				result = append([]int{index-i}, result...)
			}
		}
	}
	return result
}

func getIndex(value int, i int, m map[int]int, T []int) int {
	v, ok := m[i]; if ok {
		if v == 0 {
			return  0
		} else {
			if T[v] > value {
				return v
			} else {
				return getIndex(value, v, m, T)
			}
		}
	} else {
		return 0
	}
}

func main() {
	input:= []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(input))
}
