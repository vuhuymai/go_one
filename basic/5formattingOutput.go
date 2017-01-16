package basic

import (
	"fmt"
)

func UsingPrintf1() {
	name := "Mittens"
	weight := 12
	// Use %s to mean string.
	// ... Use an explicit newline.
	fmt.Printf("The cat is named %s.\n", name)
	// Use %d to mean integer.
	fmt.Printf("Its weight is %d.\n", weight)
}

//the Print and Printf
func UsingPrintWithForLoop()  {
    for i := 0; i < 10; i++ {
        fmt.Print(i, " ");
    }
    fmt.Println();
}

func UsingPrintfWithLoop() {
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

// to make programs simpler, we can just use the "%v" format code to insert values
// The v code handles ints, bools, strings and other values. It makes Printf calls easier to write and read
func UsingTheVCode()  {
    result := true 
    name := "Spark" 
    size := 2000
    
    // Print line with v format codes.
    fmt.Printf("Result = %v, Name = %v, Size = %v\n", result, name, size)
}
