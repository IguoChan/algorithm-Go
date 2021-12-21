package main

import (
	"fmt"
	"time"
)

var ii int

func a(i int) {
	ii += i
}

func main() {
	s := make([]int, 10000000)

	t := time.Now().UnixNano()
	l := len(s)
	for i := 0; i < l; i++ {
		a(i)
	}
	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t)
	s = append()
}
