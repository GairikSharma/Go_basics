package main

import "fmt"

func main() {

	fmt.Println("Enter the range ?")
	var n int
	fmt.Scanln(&n)
	// for i := 0; i<n; i++ {
	// 	fmt.Println("The number is ", n)
	// }
	arr := make([]rune, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the value:")
		fmt.Scanf("%c\n", &arr[i])
	}
	for index, char := range arr {
		fmt.Printf("Index: %d Value: %c\n", index, char)
	}
}
