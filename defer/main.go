package main

import "fmt"

func sum (a int, b int) int {
	return a+b
}

func main() {

	fmt.Println("Line 1")
	fmt.Println("Line 2")
	res := sum(5,2)
	defer fmt.Println("Result is: ", res)
	fmt.Println("Line 3")
}
