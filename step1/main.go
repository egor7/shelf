// Options example
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

package main

import (
	"fmt"
	"github.com/egor7/shelf/step1/pkg"
)

func main() {
	num := pkg.Newmyint(10, 5)

	sz := num.Option(pkg.Size(50))
	fmt.Println(num)

	num.Option(sz)
	fmt.Println(num)
}
