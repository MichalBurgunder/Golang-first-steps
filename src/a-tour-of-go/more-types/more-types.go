package main

import "fmt"

func main() {
	// pointers()
	// structures()
	// structsWithPointers()
	// structLiterals()
	// arrays()
	// slices()
	// moreSlices()
	// sliceLiterals()
	sliceDefaults()
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
