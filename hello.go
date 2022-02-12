package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello")

	//Variables

	var x int = 10
	fmt.Println(x)

	// Conditions
	if x > 6 {
		fmt.Println("More than 6")
	}

	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 11 {
		fmt.Println("Less than 11")
	} else {
		fmt.Println("Nothing.")
	}

	// Arrays

	var a [5]int
	fmt.Println(a)

	a[2] = 10

	fmt.Println(a)

	//short-hand Arrays
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//Slices

	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c)
	c = append(c, 10)
	fmt.Println(c)

	// Associative arrays or key-value pair called map
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	fmt.Println(vertices)

	delete(vertices, "triangle")
	fmt.Println(vertices)

	// Loop only loop is for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// make as while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// range based for loop

	for index, value := range c {
		fmt.Println("index", index, "value", value)
	}

	// range can be use with map also

	for key, value := range vertices {
		fmt.Println("key", key, "value", value)
	}

	// Custom Function

	fmt.Println(sum(10, 10))

	result, err := sqrt(-1)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Output", result)
	}

	// struct
	p := person{name: "Hassan", age: 34}
	fmt.Println(p)

	persons := []person{{name: "Hassan", age: 34}, {name: "Aaliya", age: 7}}
	fmt.Println(persons)
	fmt.Println(persons[0])
	fmt.Println(persons[0].name)

	//pointers

	k := 5
	fmt.Println(&k)
	inc(&k)
	fmt.Println(k)
}

func inc(x *int) {
	*x++
}
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
