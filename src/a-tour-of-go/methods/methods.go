package main

import (
	"fmt"
	"math"
	"time"
)

// pointers receivers()
func main() {
	// var a Abser
	// f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}
	// v2 := Vertex{30, 40}

	// a = f  // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser

	// // In the following line, v is a Vertex (not *Vertex)
	// // and does NOT implement Abser.
	// a = &v2

	// fmt.Println(a.Abs())
	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// }
	// var it ErrNegativeSqrt = 33
	fmt.Println(Sqrt(3334))
	fmt.Println(Sqrt(-209))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return "cannot Sqrt negative number"
	} else {
		num, _ := Sqrt(44.0)
		s := fmt.Sprintf("%f", num)
		return s
	}
}

func Sqrt(x float64) (float64, string) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x).Error()
	} else {
		guess := 1.0
		inc := 1.0

		for i := 0; i < 100; i++ {
			if math.Pow(guess, 2) < x {
				inc = inc / 2
				guess += inc
			} else {
				inc = inc / 2
				guess -= inc
			}
		}

		return guess, ""
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type IPAddr [4]int

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type I interface {
	M()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct {
	S string
}

type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func (t *T) M() {
// 	fmt.Println(t.S)
// }

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (v *Vertex)
// must be a pointer in order to change the value
func Scale2(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
