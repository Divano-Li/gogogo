package main

func getRow(rowIndex int) []int {
	a := make([]int, rowIndex + 1)
	getRow1(rowIndex, a)
	return a
}

func getRow1(rowIndex int, a []int){
	if rowIndex == 0 {
		a[0] = 1
		return
	}

	if rowIndex == 1 {
		a[0] = 1
		a[1] = 1
		return
	}
	getRow1(rowIndex -1, a)
	b := make([]int, len(a))
	copy(b, a)
	for i:= 1; i<= rowIndex-1; i++ {
		a[i] = b[i-1] + b[i]
	}
	a[rowIndex] = 1

}

func getRow3(rowIndex int) []int {
	a := make([]int, rowIndex + 1)
	if rowIndex == 0 {
		a[0] = 1
		return a
	}
	if rowIndex == 1 {
		a[0] = 1
		a[1] = 1
		return a
	}
	b := make([]int, rowIndex + 1)
	for i:=1; i<= rowIndex ; i++ {

		for j:=0 ; j<= i; j++ {
			if i==j || j == 0 {
				b[j] = 1
			} else {
				b[j] = a[j-1] + a[j]
			}
		}
		copy(a, b)
	}
	return a
}
