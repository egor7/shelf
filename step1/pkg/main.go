package pkg

import (
	"fmt"
)

type myint struct {
	n int

	size int // options
}

func Newmyint(pn int, psize int) myint {
	var num myint

	num.n = 10
	num.size = 5

	return num
}

func (m myint) String() string {
	return fmt.Sprintf("%d(%d)", m.n, m.size)
}

type optfn func(*myint) optfn

func (m *myint) Option(f optfn) (oldfn optfn) {
	return f(m)
}

func Size(sz int) optfn {
	return func(i *myint) optfn {
		prev := i.size
		i.size = sz
		return Size(prev)
	}
}
