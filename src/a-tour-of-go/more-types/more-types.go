package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tour/pic"
)

func main() {
	// pointers()
	// structures()
	// structsWithPointers()
	// structLiterals()
	// arrays()
	// slices()
	// moreSlices()
	// sliceLiterals()
	// sliceDefaults()
	// slicelengthCapacity()
	// nilSlices()
	// makeFunction()
	// nestedSlices()
	// appendingSlices()
	// tryingOutRange()
	// rangeContinued()
	// exerciseFunction()
	// maps()
	// mutationMaps()
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))
	// functionClosure()
	// wc.Test(WordCount)
	// exerciseMaps("good morning Santa Barbara! I hope you are doing well")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	num0 := 0
	num1 := 1
	return func() int {
		if num0 > num1 {
			num1 = num0 + num1
			return num1
		} else {
			num0 = num0 + num1
			return num0
		}

	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func functionClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func WordCount(s string) map[string]int {
	// var word string = ""
	// var inc int = 1
	// var finalMap map[string]int

	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
	// for l := 0; l < len(s); l++ {

	// 	if string(s[l]) == " " {
	// 		finalMap[word] = inc
	// 		word = ""
	// 		inc++
	// 	} else {
	// 		word = word + string(s[l])
	// 	}
	// }
	// return finalMap
}
func mutationMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

type Vertex2 struct {
	Lat, Long float64
}

func maps() {
	var m map[string]Vertex2
	var m2 map[int64]Vertex2
	fmt.Print(m2)
	m = map[string]Vertex2{
		"Bell Labs":     {40.68433, -74.39967},
		"some location": {40.0, -74.4},
	}

	fmt.Println(m)
}

func Pic(dx, dy int) [][]uint8 {
	finalArray := make([][]uint8, dy)

	for i := 0; i < dx; i++ {
		finalArray[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			finalArray[i][j] = 1
			switch {
			case j*i < 1000:
				finalArray[i][j] = 0
			case j*i < 2000:
				finalArray[i][j] = 150
			// case j%5 == 0:
			// 	finalArray[i][j] = 200
			default:
				finalArray[i][j] = 255
			}
		}
	}
	return finalArray
}

func exerciseFunction() {
	pic.Show(Pic)
}
func rangeContinued() {
	var t int64 = 2
	var z int64 = 1 << uint(t)
	var r int64 = z << uint(1)
	fmt.Println("z")
	fmt.Println(z)
	fmt.Println(r)
	fmt.Printf(strconv.FormatInt(r, 2))
	fmt.Println(r)
	pow := make([]int, 10)
	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func tryingOutRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func appendingSlices() {
	var s []int
	printSlice2(s)

	s = append(s, 0)
	printSlice2(s)

	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func nestedSlices() {
	board := [][]string{[]string{"-", "-", "-"}, []string{"-", "-", "-"}, []string{"-", "-", "-"}}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func makeFunction() {
	a := make([]int, 0, 5)
	b := []int{1, 2, 3, 4, 5}
	c := b[:2]
	d := c[2:5]
	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)
}

func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice2(slice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}

func printSlice(str string, slice []int) {
	// fmt.Println("len%d cap%d %v\n", len(s), cap(s))
	fmt.Printf("%s len=%d cap=%d %v\n",
		str, len(slice), cap(slice), slice)
}

func slicelengthCapacity() {
	s := []int{2, 3, 5, 7, 11}
	printSlice("s", s)
}
func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	fmt.Println(s)

	t := s[:2]
	fmt.Println(t)
	fmt.Println("My test")
	z := s[1:2]
	fmt.Println(z)

	r := s[1:]
	fmt.Println(r)
	fmt.Println("s")
	fmt.Println(s)
	fmt.Println("My test")
	o := s[1:4]
	fmt.Println(o)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	fmt.Println(q)
	r := []bool{true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}

func pointers() string {
	i, j := 42, 2701
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(j)
	return "OK!"
}

type Vertex struct {
	Y int
	X int
}

func structures() {

	fmt.Println(Vertex{22, 66})
	fmt.Println(Vertex{7, 9}.X)
}

func structsWithPointers() {
	v := Vertex{7, 9}
	fmt.Println(v)
	p := &v
	fmt.Println(&p.X)
	p.X = 1e3
	fmt.Println(v)
}

func structLiterals() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}
	fmt.Println(v1, v2, v3, p)
}

func arrays() {
	var a []string
	a[0] = "Hello"
	a[1] = "World"
	a[6] = "how u do"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[3:6])
}

func moreSlices() {

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	} // the Beatles
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
