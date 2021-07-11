package main

import (
	"strconv"
	"strings"
)

func openLock(deadends []string, target string) int {
	if "0000" == target {
		return 0
	}
	start := "0000"
	count := -1
	used:= map[string]int{}
	for _,v := range deadends {
		used[v] = 1
	}
	_, ok1 := used[start]; if ok1{
		return -1
	}
	queue := ConstructQueue(10000)
	queue.enQueue(start)
	key := ""
	i1 := "0"
	i2 := "0"
	i3 := "0"
	i4 := "0"
	i := 0
	k := 0
	temp := make([]string, 4)
	for queue.isEmpty() == false {
		count++
		times := queue.tail - queue.head + 1
		for jj:=0; jj< times; jj++ {
			key = queue.front()
			queue.deQueue()
			temp = strings.Split(key, "")
			i1 = temp[0]
			i2 = temp[1]
			i3 = temp[2]
			i4 = temp[3]
			for j:=0; j<len(temp); j++ {
				i, _ = strconv.Atoi(temp[j])
				k = i + 1
				if k == 10 {
					k = 0
				}
				if j ==0 {
					key = strconv.Itoa(k)+i2+i3+i4
				} else if j==1 {
					key = i1+strconv.Itoa(k)+i3+i4
				} else if j==2 {
					key = i1+i2+strconv.Itoa(k)+i4
				} else if j==3 {
					key = i1+i2+i3 + strconv.Itoa(k)
				}

				_, ok := used[key]; if !ok {
					if key == target {
						return count+1
					}
					queue.enQueue(key)
					used[key] = 1
				}

				k = i-1
				if k == -1 {
					k =9
				}
				if j ==0 {
					key = strconv.Itoa(k)+i2+i3+i4
				} else if j==1 {
					key = i1+strconv.Itoa(k)+i3+i4
				} else if j==2 {
					key = i1+i2+strconv.Itoa(k)+i4
				} else if j==3 {
					key = i1+i2+i3 + strconv.Itoa(k)
				}
				_, ok = used[key]; if !ok {
					if key == target {
						return count+1
					}
					queue.enQueue(key)
					used[key] = 1

				}
			}
		}
	}
    return  -1
}