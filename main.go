package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator
func gen3(tag string, c chan<- string) {
	for i := 0; ; i++ {
		ms := rand.Intn(1000)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		c <- fmt.Sprintf("		%s,hello.%d:%d from gen3()", tag, i, ms)
	}
}

func gen2(tag string) <-chan string {
	fmt.Println("	", tag, "/beg")

	c := make(chan string)
	go gen3(tag, c)

	fmt.Println("	", tag, "gen2/end")
	return c
}

func flop(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func main() {
	fmt.Println("main/beg")

	c1 := gen2("c1")
	c2 := gen2("c2")
	c := flop(c1, c2)

	for i := 0; i < 20; i++ {
		//fmt.Println(<-c2)
		//fmt.Println(<-c1)
		fmt.Println(<-c)
	}

	fmt.Println("main/end")
}
