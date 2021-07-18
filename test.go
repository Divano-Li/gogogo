package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(1)
	fmt.Println([]byte("b")[:1])
	buf := make([]byte, 1024)
	f, _ := os.Open("./test.go")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}
