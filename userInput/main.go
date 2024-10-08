package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("What is your name ?")

	// var name string
	// fmt.Scan(&name) // Scan method only takes input untill it reaches a whitespace
	// fmt.Scanf("%s\t", &name)
	// fmt.Scanln(&name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name) 

}
