package main

import (
	"fmt"
)

func userInput1()  {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println(output);
}