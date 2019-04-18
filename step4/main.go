package main

import (
	"fmt"
)

func f(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	fmt.Println("=== beg")
	f(10, 20, 30)
	fmt.Println("=== end")
}
