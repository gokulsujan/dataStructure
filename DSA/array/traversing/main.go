package main

import "fmt"

func main() {
	var limit int
	fmt.Println("Enter the array size")
	fmt.Scan(&limit)
	arr := make([]int, limit)
	fmt.Println("Enter the values of array")
	for i := 0; i < limit; i++ {
		fmt.Println("arr[", i, "]: ")
		fmt.Scan(&arr[i])
	}
	fmt.Println("The array elements are: ", arr)

	fmt.Println("Printing the element one by one")
	for i := 0; i < limit; i++ {
		fmt.Println(arr[i])
	}

}
