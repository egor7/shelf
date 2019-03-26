package main

import (
	"fmt"
//1	"math/rand"
//1	"time"
)

//1 // Generator
//1 func gen3(tag string, c chan<- string) {
//1 	for i := 0; ; i++ {
//1 		ms := rand.Intn(1000)
//1 		time.Sleep(time.Duration(ms) * time.Millisecond)
//1 		c <- fmt.Sprintf("		%s,hello.%d:%d from gen3()", tag, i, ms)
//1 	}
//1 }
//1 
//1 func gen2(tag string) <-chan string {
//1 	fmt.Println("	", tag, "/beg")
//1 
//1 	c := make(chan string)
//1 	go gen3(tag, c)
//1 
//1 	fmt.Println("	", tag, "gen2/end")
//1 	return c
//1 }
//1 
//1 func flop(c1, c2 <-chan string) <-chan string {
//1 	c := make(chan string)
//1 	go func() {
//1 		for {
//1 			c <- <-c1
//1 		}
//1 	}()
//1 	go func() {
//1 		for {
//1 			c <- <-c2
//1 		}
//1 	}()
//1 	return c
//1 }
//1 
//1 func main() {
//1 	fmt.Println("main/beg")
//1 
//1 	c1 := gen2("c1")
//1 	c2 := gen2("c2")
//1 	c := flop(c1, c2)
//1 
//1 	for i := 0; i < 20; i++ {
//1 		//fmt.Println(<-c2)
//1 		//fmt.Println(<-c1)
//1 		fmt.Println(<-c)
//1 	}
//1 
//1 	fmt.Println("main/end")
//1 }

// Options example
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type circle struct {
	x float64
	y float64
	R float64

	v int // verbosity
}
func NewCircle(x, y, R float64) circle {
	c := circle{}
	c.x = x
	c.y = y
	c.R = R
	return c
}

func (c circle) String() string {
	if c.v != 0 {
		return fmt.Sprintf("(%v,%v):%v", c.x, c.y, c.R)
	} else {
		return fmt.Sprintf("(%v,%v)", c.x, c.y)
	}
}

type option func(*circle) option

func (c *circle) Option(opts... option) (prev option) {
	for _, opt := range opts {
		prev = opt(c)
	}
	return prev
}

func Verbosity(v int) option {
	return func(c *circle) option {
		prev := c.v
		c.v = v
		return Verbosity(prev)
	}
}

func main () {
	c := NewCircle(1,2,3)

	v := c.Option(Verbosity(1))
	c.Option(Verbosity(1))
	fmt.Println(c)

	c.Option(v)
	fmt.Println(c)
}
