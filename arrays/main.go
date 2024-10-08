package main

import "fmt"

func main() {
	fmt.Println("Enter the size of the array (minimum length is 3): ")
	var size int
	fmt.Scan(&size)


	// arr := make([]int, size)
	// arr[1]=2
	// if len(arr) > 3 {
	// 	arr[3]=5
	// }
	// fmt.Println(arr)
	// fmt.Println("Capacity: ", cap(arr))

	//Now we are going to make slices

	
	arr := []int{}
	for i := 0; i < size; i++ {
		var element int
		fmt.Println("Enter the element at index", i)
		fmt.Scan(&element)
		arr = append(arr, element)
	}
	fmt.Println(arr)

}
