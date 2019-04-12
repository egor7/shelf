package pkg

import (
	"fmt"
)

type I interface {
	M()
}

func NewI() I {
	return newS1()
}

type s1 struct {
	n  int
	sz int
}

func (_ s1) M() {
	fmt.Println("in s1.M")
}
func (_ s1) N() {
	fmt.Println("in s1.N")
}

func newS1() *s1 {
	return &s1{
		n:  99,
		sz: 2,
	}
}

type S2 struct {
	n  int
	sz int
}

func (_ S2) M() {
	fmt.Println("in S2.M")
}
