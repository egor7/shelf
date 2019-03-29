// Options example
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

package main

import (
	"fmt"
)

type myint struct {
	n    int
	size int
}

func (m myint) String() string {
	return fmt.Sprintf("%d(%d)", m.n, m.size)
}

type fretf func(*myint) fretf

func (m *myint) Option(newfn fretf) (oldfn fretf) {
	return newfn(m)
}

func Optsize(sz int) fretf {
	return func(i *myint) fretf {
		prev := i.size
		i.size = sz
		return Optsize(prev)
	}
}

func main() {
	var num myint
	num.n = 10
	num.size = 5

	osz := num.Option(Optsize(50))
	fmt.Println(num)

	osz2 := num.Option(osz)
	fmt.Println(num)

	osz3 := num.Option(osz2)
	fmt.Println(num)

	num.Option(osz3)
	fmt.Println(num)
}
