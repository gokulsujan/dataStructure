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
	var index int
	fmt.Println("Enter the array index you want to update")
	fmt.Scan(&index)
	if index >= limit {
		fmt.Println("Sorry the index you entered is the more than the limit")
	} else {
		var val int
		fmt.Println("Enter the value you want to update")
		fmt.Scan(&val)
		arr[index] = val
	}
	fmt.Println(arr)

}
