package main

import "fmt"

func initArray() {
	var x [5]int
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	fmt.Println(x) //output: [0 0 0 0 100]

	//a shorter way to create array
	y := [5]float64{23.5, 33.1, 16.8, 133.113, -88.17}
	fmt.Println(y)

	//can initialize the array which is too long over multiple lines
	z := [5]float64{
		98.234,
		93.00,
		77.66,
		82.199,
		-83.177, // ==> notice the extra trailing here, required by Go
	}
	fmt.Println(z)
}

func usingArrLen() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
	//without type conversion, there will be an error
	//to convert between types, you use the type name like a function
}

func theRangeKeyword() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	//use the keyword range followed by the name of the variable we want to loop over
	for i, value := range x {
		fmt.Println("Element", value, "is at position", i, "in the array.")
	}
	fmt.Println()

	//if we don't want the index, we can ommit it by using the underscore _
	for _, value := range x {
		fmt.Print(value, " ")
	}
	fmt.Println()

}

func initSimpleSlice() {
	x := make([]float64, 5)
	fmt.Println(x) // output: [0 0 0 0 0]

	//the third parameter (array capacity) of the make function
	y := make([]float64, 5, 10) //A slice of length 5 with a capacity of 10
	fmt.Println(y)              // output: [0 0 0 0 0]

	//using the [low:high] expression to make a slice from an array
	//low: where to start the slice
	//high: where to end but NOT including
	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[0:len(arr)]
	fmt.Println(z) // output: [1 2 3 4 5]
	f := arr[1:4]
	fmt.Println(f) // output: [2 3 4]

	/*	allowed to omit low, high, or even both low and high
		arr[0:] is the same as arr[0:len(arr)]
		arr[:5] is the same as arr[0:5]
		arr[:] is the same as arr[0:len(arr)]
	*/
}

func basicSliceOperations() {

	//indexing like working with array

	//append function
	fmt.Println("Append function")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2) //output [1,2,3,4,5]

	//copy function
	fmt.Println("Copy function")
	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = make([]int, 3) // 3 means the capacity of the underlying array of slice 2
	copy(slice2, slice1) //only the first 3 elements of slice1 are copied
	fmt.Println(slice2) // output: [1, 2, 3]
}
