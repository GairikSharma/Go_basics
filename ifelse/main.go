package main

import "fmt"

func main() {

	fmt.Println("Enter your age ?")
	var num int
	fmt.Scan(&num)
	if num >= 18 {
		fmt.Println("You are eligible for voting")
	} else {
		fmt.Println("You are not eligible for voting")
	}

}
