package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func printNumbers(n int) []int {
	max := math.Pow10(20) -1
	 var i float64 = 1
	 result := make([]int, 0)
	for ; i< max; i++ {
      result = append(result, int(i))
	}
	fmt.Println(max)
	return []int{}
}


