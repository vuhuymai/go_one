package main

import "fmt"

import B "github.com/vuhuymai/vuone/basic"

func main() {

	fmt.Println(B.VariadicAddFunc())
	fmt.Println(B.VariadicAddFunc(1,2,3,4,5))
	xs := []int{3, 7, 87, 10, -230, 55}
	fmt.Println(B.VariadicAddFunc(xs...))}

func variable1() {
	var s = "Hello variable1"
	fmt.Println(s)

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}
