// Options example
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

package main

import (
	"github.com/egor7/shelf/step2/pkg"
)

func main() {
	var t = pkg.NewI()
	t.M()

	//var t pkg.S2
	//t.M() // in S2.M

	//var t pkg.S2
	//t = *pkg.NewI().(*pkg.S2) // interface conversion: pkg.I is *pkg.s1, not *pkg.S2
	//t.M()
}
