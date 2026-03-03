package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

func areaOfCircle(radius float64, pi float64) float64 {
	return radius * radius * pi
}

func main() {
	var i int
	i = 10

	var s string = "string"

	f := true

	var x, y float64 = 1.3, 1.8

	a, b := 1.2, 1.5

	z := cmplx.Sqrt(complex(-a, b))

	var complNum float64 = float64(math.Sqrt((float64(x*x + y*y)))) // type conversion

	const (
		Pi     = 3.14
		radius = 5
		Big    = 1 << 100
		Small  = Big >> 99
	)

	fmt.Printf("Type is %T, and value is %v, and string is %q\n", i, i, s)
	fmt.Println(f, z, complNum)
	fmt.Printf("Area of circle of radius %v is %v\n", radius, areaOfCircle(radius, Pi))
	fmt.Println(float64(Big), Small)

	// loops
	var sum int = 0

	for i = 0; i < 100; i++ { // for loop
		sum += i
	}

	itr := 0
	for itr < 100 { // while loop
		sum += itr
		itr++
	}

	fmt.Println("Sum:- ", sum)

	// if-else
	if sum += 100; sum < 10000 {
		fmt.Println("less then 10k")
	} else if sum == 10000 {
		fmt.Println("equal to 10k")
	} else {
		fmt.Println("greater then 10k")
	}

	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// defer
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%v ", i) // push onto a stack    after below code done executing then 2 1 0
	}
	fmt.Println("defer")

}
