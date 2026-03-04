package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	// "math"
	// "math/cmplx"
	// "runtime"
)

type Domain int

const (
	FRONTEND Domain = iota // FRONTEND is assigned 0
	BACKEND                // BACKEND is assigned 1
	DEVOPS
	MOBILE
	UIUX
)

// Implement the Stringer interface for the Day type
func (d Domain) String() string {
	return []string{"FRONTEND", "BACKEND", "DEVOPS", "MOBILE", "UIUX"}[d]
}

type Employee struct {
	name     string
	ID       string
	domainId Domain
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee:- name: %s, ID: %s, domainId: %v", e.name, e.ID, e.domainId)
}

// custom validation, this is like a member function of Employee struct
func (e *Employee) Validate() error {
	if e.ID == "" {
		return errors.New("employee ID cannot be empty")
	}
	if len(e.name) < 3 {
		return errors.New("employee name must have more then 3 characters")
	}
	return nil
}

// // 1. Define your custom type (usually a struct)
// type MyError struct {
//     Code    int
//     Message string
// }

// // 2. Implement the Error() method (satisfies the 'error' interface)
// // Using a pointer receiver is standard practice for errors
// func (e *MyError) Error() string {
//     return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
// }

// func doSomething() error {
//     // Return your custom error as the 'error' interface type
//     return &MyError{Code: 404, Message: "Not Found"}
// }

// pointer vs value reciever

func sliceInfo(someSlice []struct {
	key   string
	value uint
}) {
	fmt.Println(someSlice, "of length", len(someSlice), "and capacity of", cap(someSlice))
}

func checkNil(someSlice []struct {
	key   string
	value uint
}) {
	if someSlice != nil {
		fmt.Println("slice is not NIL")
	} else {
		fmt.Println("slice is NIL")
	}
}

func apply(fn func(float64, float64) float64, p1 float64, p2 float64) float64 {
	return fn(p1, p2)
}

func compute(option string) func(int) int {
	var value int
	if option == "adder" {
		value = 0
		return func(x int) int {
			value += x
			return value
		}
	} else if option == "multiplier" {
		value = 1
		return func(x int) int {
			value *= x
			return value
		}
	} else {
		value = 0
		return func(x int) int {
			return value
		}
	}
}

func main() {
	fmt.Println("------------- Struct ----------------")
	var (
		e1 = Employee{name: "emp-1-backend-1", ID: "b-1", domainId: BACKEND}
		e2 = &Employee{name: "emp-2-frontend-1", ID: "f-1", domainId: FRONTEND}
	)

	if err := e1.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Employee is valid")
	}

	fmt.Println(e1)
	fmt.Println(e2)

	// err := doSomething()
	// if err != nil {
	// 	fmt.Println(err) // Automatically calls Error() method
	// }

	fmt.Println("------------- Array ----------------")

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	evenIntegers := [5]int{2, 4, 6, 8, 10}
	fmt.Println(evenIntegers)

	fmt.Println("------------- Slice ----------------")
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.

	mp := []struct { // array of struct of string and int
		key   string
		value uint
	}{
		{"apple", 2},
		{"mango", 5},
		{"orange", 6},
	}

	// slice of array
	slice1 := mp[0:1]
	slice2 := mp[:2]
	slice3 := mp[1:]
	sliceEmpty := mp[3:]

	// update of array
	mp[1] = struct {
		key   string
		value uint
	}{"pineapple", 2}

	fmt.Println("Array: ", mp)
	sliceInfo(slice1)
	sliceInfo(slice2)
	sliceInfo(slice3)
	sliceInfo(sliceEmpty)
	checkNil(sliceEmpty)

	var arrayEmpty []struct {
		key   string
		value uint
	}
	sliceInfo(arrayEmpty)
	checkNil(arrayEmpty) // why sliceEmpty is not NIL and arrayEmpty is NIL

	tictactoe := [][]string{
		{"O", "x", "O"},
		[]string{"x", "O", "O"},
		{"O", "x", "x"},
	}
	for i := range len(tictactoe) {
		fmt.Printf("%s\n", strings.Join(tictactoe[i], " "))
	}

	mp = append(mp, struct {
		key   string
		value uint
	}{"kiwi", 1})

	for i, v := range mp { // iteration on array
		fmt.Println(i, v.key, v.value)
	}

	arrayUsingMake := make([]int, 5, 10) // length 5 and capacity 10
	fmt.Println(arrayUsingMake, "of length", len(arrayUsingMake), "and capacity of", cap(arrayUsingMake))

	fmt.Println("--------------- Map --------------------")

	// map using make
	var m map[string]Employee
	m = make(map[string]Employee)
	m["e1"] = Employee{name: "emp-1-backend-1", ID: "b-1", domainId: BACKEND}
	fmt.Println("map using make:- ", m)

	// map using literals
	var mpp = map[string]Employee{
		"e1": Employee{name: "emp-1-backend-1", ID: "b-1", domainId: BACKEND},
		"e2": {name: "emp-2-frontend-1", ID: "f-1", domainId: FRONTEND},
	}
	fmt.Println("map using literals:- ", mpp)

	// map operations
	mpp1 := make(map[string]int)
	mpp1["1"] = 1
	mpp1["1"] = 2
	delete(mpp1, "1")
	v, ok := mpp1["1"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("------------- Functions ----------------")

	// higer order function
	customDistanceFunction := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(apply(customDistanceFunction, 3, 4))
	fmt.Println(apply(math.Pow, 3, 4))

	// closure
	adder := compute("adder")
	for i := 1; i < 10; i++ {
		fmt.Printf("%v ", adder(i))
	}
	fmt.Println("-> value are like this because of closure")

	// recievers:- value, pointer

	fmt.Println("------------- Interface ----------------")
}
