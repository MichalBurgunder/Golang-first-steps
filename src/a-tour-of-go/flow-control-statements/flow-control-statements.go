package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// func main() {
// 	const name, age = "Kim", 22
// 	s := fmt.Sprint(name, " is ", age, " years old.\n")

// 	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

// }

func main() {
	fmt.Println("Here we go:")
	// forLoop(8876)
	// goesOnAndOnAndOnAndOn()
	// fmt.Println(simpleIf())
	// x := 55.08
	// simpleSwitch()
	// usingTime()
	// fmt.Println(deferring())
	stackedDefers()
	// anotherString := fmt.Sprint(tryingOutSprint("Michal", 27))
	// t := fmt.Sprint(math.Sqrt(x))
	// fmt.Println(t)
	// pow(2, 20, 100000)
}

func stackedDefers() string {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	return "OK!"
}
func forLoop(number int) int {
	sum := 1
	for sum < number {
		sum += sum
		fmt.Println(sum)
	}

	return sum
}

func goesOnAndOnAndOnAndOn() {
	for {
	}
	// fmt.Println("No sir, we haven't seen him")
}

func deferring() string {
	defer fmt.Println("ready")
	fmt.Println("I'm")
	return "I'm done!"
}

func simpleIf() int {
	t := "lalala"
	z := 59

	if t == "alalal" {
		return 44
	} else if z == 56 {
		return 56
	} else {
		return 5
	}
}

func tryingOutSprint(name string, age int) string {
	s := fmt.Sprint(name, " is ", age, " years old.\n")
	return s
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%e>=%g\n", v, lim)
	}
	return lim
}

func customSqrt(number float64) float64 {
	finalNumber := 1.0

	return finalNumber
}

func simpleSwitch() {
	os2 := runtime.GOOS
	fmt.Println(fmt.Printf("My OS is %v", os2))
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func usingTime() {

	today := time.Now().Weekday()
	switchCase := "greatScott"
	if switchCase == "switchySwitch" {
		fmt.Println("When is Saturday?")
		switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow")
		case today + 2:
			fmt.Println("in two days")
		default:
			fmt.Println("too far")
		}
	} else if switchCase == "theTimeWillCome" {
		// myLocation, someError := time.LoadLocation("Rheinfelden")
		// fmt.Println(time.Now().Add())
		t1 := time.Now()
		// fmt.Println("hi")
		t2 := time.Now()
		elapsed := t2.Sub(t1)
		fmt.Println(elapsed)
		// 	newDate := time.Date(2019, 12, 6, 4, 20, 6, 0, myLocation)
		// 	fmt.Println(fmt.Sprint(newDate))
		// fmt.Println(stringifiedAfter)
	} else if switchCase == "greatScott" {
		start := time.Now()
		// ... operation that takes 20 milliseconds ...
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
	}
}
