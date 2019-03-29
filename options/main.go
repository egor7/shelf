// Options example
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

package main

import (
	"fmt"
)

type circle struct {
	x float64
	y float64
	r float64

	v int // verbosity
}

func NewCircle(x, y, r float64) circle {
	c := circle{}
	c.x = x
	c.y = y
	c.r = r
	return c
}

func (c circle) String() string {
	if c.v != 0 {
		return fmt.Sprintf("(%v,%v):%v", c.x, c.y, c.r)
	} else {
		return fmt.Sprintf("(%v,%v)", c.x, c.y)
	}
}

type option func(*circle) option

func (c *circle) Option(opts ...option) (prev option) {
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

func main() {
	c := NewCircle(1, 2, 3)

	v := c.Option(Verbosity(1))
	fmt.Println(c)

	c.Option(v)
	fmt.Println(c)
}
