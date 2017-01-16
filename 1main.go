package main

import "fmt"

func main() {
	checkingMapLookup()
}

func variable1() {
	var s = "Hello variable1"
	fmt.Println(s)

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}
