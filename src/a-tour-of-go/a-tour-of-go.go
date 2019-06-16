package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool
var a int8 = 127
var b uint8 = 55
var complex complex128 = -22 + 77i

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var i int
	// c := true
	fmt.Println(a)
	fmt.Println(i, c, python, java)
	fmt.Println(cmplx.Sqrt(complex))
	fmt.Printf("Type: %T Value: %v\n", complex, complex)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}
