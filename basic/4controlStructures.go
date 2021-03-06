package basic

import (
	"fmt"
)

func SimpleForLoop() {
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

func FullForLoop() {
    for index := 0; index < 10; index++ {
        fmt.Printf("%v ", index)
    }
    fmt.Println();
}

func SimpleIfElse()  {
    for index := 1; index <= 10; index++ {
        if index % 2 == 0 {
            fmt.Println(index, "is even");
        } else {
            fmt.Println(index, "is odd");
        }
    }
}

func SimpleSwitch()  {
    for i := 1; i <= 10; i++ {        
        switch i {
        case 1:
            fmt.Println("one");
        case 2:
            fmt.Println("two");
        case 3:
            fmt.Println("three");
        default: fmt.Println("other number");
        }
    }
}
