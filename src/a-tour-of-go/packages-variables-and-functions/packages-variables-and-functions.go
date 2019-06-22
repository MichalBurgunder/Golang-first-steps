package main

import (
	"fmt"
	"math"
	"reflect"
)

var c, python, java bool
var a int8 = 127
var b uint8 = 55
var complex complex128 = -22 + 77i

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// var i int
	// c := true
	// CosaSpeciale()
	// apprendimento()
	// workingWithArrays()
	// chinAndConst()
	fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
	fmt.Println(binaryFun(53))
	// var complex complex128 = 55 - 14i
	// fmt.Println(i, c, python, java)
	// fmt.Println(cmplx.Sqrt(complex))
	// fmt.Printf("Type: %T Value: %v\n", complex, complex)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func CosaSpeciale() string {
	var (
	// i int = 777
	// f float64
	// b bool
	// s string = "lalala"
	)
	var x, y int = 33, 4
	var z float64 = math.Sqrt(float64(x*x + y*y))
	// fmt.Println("%v %v %v %q\n", x, f, b, s)
	fmt.Println(x, y, z)
	return "finito!"
}

func apprendimento() string {
	// var i int
	// j := i // j is an int
	var (
		i   int     = 777
		fl  float64 = 1.53
		str string  = "boop"
		sli [4]string
		boo bool = true
	)
	sli[0] = "hi"
	sli[1] = "there"
	sli[2] = "you"
	sli[3] = "!"
	fmt.Printf("type: %T\n", i)
	fmt.Printf("boolean: %t\n", boo)
	fmt.Printf("Base8: %o\n", i)
	fmt.Printf("scientific notation: %e\n", fl)
	fmt.Printf("base64: %X\n", str)
	fmt.Printf("just the slice: %v\n", sli)
	return "We made it!"
}

func workingWithArrays() string {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2)
	printSlice(s)

	s = append(s, 777, 888)
	printSlice(s)

	s = append(s, 112, 221, 667, 98989)
	printSlice(s)

	s = s[8:]
	printSlice(s)

	return "OK!"
}

func chinAndConst() {
	const World = "世界"
	const Truth = true

	// fmt.Println("%T", Truth)
	fmt.Println(reflect.TypeOf(Truth))
	fmt.Println(reflect.TypeOf(World))
	fmt.Println("Go rules?", Truth)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	fmt.Println("needInt:", x)
	return x*10 + 1
}
func needFloat(x float64) float64 {
	fmt.Println("needFloat:", x)
	return x * 0.1
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func binaryFun(number int) int {
	finalNumber := number << 1
	return finalNumber
}
