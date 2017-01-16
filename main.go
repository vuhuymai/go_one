package main

import "fmt"

import B "github.com/vuhuymai/vuone/basic"

func main() {

	fmt.Println(B.VariadicAddFunc())
	fmt.Println(B.VariadicAddFunc(1, 2, 3, 4, 5))
	xs := []int{3, 7, 87, 10, -230, 55}
	fmt.Println(B.VariadicAddFunc(xs...))
}
