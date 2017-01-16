package basic

import "fmt"

func InitArray() {
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

func UsingArrLen() {
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

func TheRangeKeyword() {
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

func InitSimpleSlice() {

	//formal way to declare a slice
	fmt.Println("Formal way to declare a slice")
	var simpleSlice []int
	fmt.Println(simpleSlice)

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

func BasicSliceOperations() {

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
	copy(slice2, slice1)    //only the first 3 elements of slice1 are copied
	fmt.Println(slice2)     // output: [1, 2, 3]
}

func SimpleMapInitAndLookup() {

	fmt.Println("\nFormal way to declare a map")
	var x map[string]int
	fmt.Println(x)

	fmt.Println("\nMaps have to be initialized (using make function) before they can be used:")
	fmt.Println("otherwise we will get a runtime error")
	y := make(map[string]int)
	y["key1"] = 10
	z := make(map[string]string)
	z["name1"] = "Iris Duong"
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("\nRetrieve an element using the correct key:")
	fmt.Println(y["key1"])
	fmt.Println(z["name1"])

	fmt.Println("\nWhen an incorrect key is used:")
	fmt.Println("If the map contains integer elements, 0 is returned")
	fmt.Println(y["key2"])
	fmt.Println("If the map contains string elements, an empty string is returned")
	fmt.Println(z["name2"])
	fmt.Println("Technically, a map returns the zero value for the value type\n",
		"(which for strings is the empty string)")

	fmt.Println("\nAnother way to initialize a map")
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
	}
	fmt.Println(elements)
	element := elements["He"];
	fmt.Println("The element \"" + element["name"] + "\" is in \"" + element["state"] + "\" state")

}

func CheckingMapLookup() {
	names := make(map[string]string)
	names["name1"] = "Henry"
	names["name2"] = "Tom"
	names["name3"] = "Peter"

	fmt.Println("\nAccessing an element of a map can return two values instead of just one")

	value, ok := names["name1"]
	fmt.Println(value, ok) //output: Henry true

	value, ok = names["wrongKey"]
	fmt.Println(value, ok) //output: empty string "", false

	fmt.Println("\nA special way in Go to check if an element corresponding \n" +
		"to a given key is in a map")

	//this is just a closure (function)	that check if a key/value pair exists in a given map
	//Go doesn't have nested function, only function literal (i.e. closure) like this
	checkIfKeyExists := func(theMap map[string]string, givenKey string) {
		if val, yes := theMap[givenKey]; yes {
			fmt.Println("The key \""+givenKey+
				"\" exits in the map with the corresponding value:", val)
		} else {
			fmt.Println("The key \"" + givenKey +
				"\" doesn't exist in the map, no corresponding value")
		}
	}

	checkIfKeyExists(names, "name1")
	checkIfKeyExists(names, "wrongKey")
}

func BasicMapOperations() {
	//https://blog.golang.org/go-maps-in-action

}
