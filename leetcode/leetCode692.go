package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	a := map[string]int{}
	for i := 0; i < len(words); i++ {
		a[words[i]]++
	}
	uniqueString := make([]string, len(a))
	for key, _ := range a {
		uniqueString = append(uniqueString, key)
	}
	sort.Slice(uniqueString, func(i int, j int) bool {
		return a[uniqueString[i]] > a[uniqueString[j]] || (a[uniqueString[i]] == a[uniqueString[j]] && uniqueString[i] < uniqueString[j])
	})

	return uniqueString[0:k]
}

func main() {
	a := map[string]int{}
	a["a"] = 1
	a["b"] = 2
	c := []string{"a", "b", "c", "d", "a", "b"}
	fmt.Println(topKFrequent(c, 1))
	fmt.Println("abb" < "abc")
}
