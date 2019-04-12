package main

import (
	"fmt"
)

type Base interface {
	f() string
}

type A struct {
	Base
	len string
}

func (a A) f() string {
	return a.len
}

func p(pb Base) {
	fmt.Println(pb.f())
}

func main() {
	var a = &A{len: "12345"}
	p(a)
}
