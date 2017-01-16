package basic

//a typical function with input parameters and return
func Average(xs []float64) float64  {
    total := 0.0
    for _, value := range xs {
        total += value
    }
    return total / float64(len(xs))
}

//function return type can have name
func NameReturnType() (r int) { // must have parentheses (round brackets) around (r int)
	r = 1
	return
}

//Multiple values can be returned
func ReturnMultipleValues() (int, int) {
    return 5, 6
    /* 
    func main()  {
        x, y := f()
    }
    */
}

/*
Multiple values are often used to return an error value along with the result 
(x, err := f()), or a boolean to indicate success (x, ok := f()).
*/

//Variadic function
func VariadicAddFunc(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}