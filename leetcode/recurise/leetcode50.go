package main

func myPow(x float64, n int) float64 {
	if n==0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1/x
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
	   return	temp * temp
	} else if n%2==1 {
	   return 	x * temp * temp
	} else {
		return 1/x * temp * temp
	}
}
