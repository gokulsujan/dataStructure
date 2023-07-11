package main

import "fmt"

func main() {
	var limit1, limit2 int
	fmt.Println("Enter the array1 size")
	fmt.Scan(&limit1)
	fmt.Println("Enter the array2 size")
	fmt.Scan(&limit2)
	arr1, arr2 := make([]int, limit1), make([]int, limit2)
	fmt.Println("Enter the values of array1")
	for i := 0; i < limit1; i++ {
		fmt.Println("arr1[", i, "]: ")
		fmt.Scan(&arr1[i])
	}
	fmt.Println("The array1 elements are: ", arr1)

	fmt.Println("Enter the values of array2")
	for i := 0; i < limit2; i++ {
		fmt.Println("arr2[", i, "]: ")
		fmt.Scan(&arr2[i])
	}
	fmt.Println("The array2 elements are: ", arr2)
	mergedArr := append(arr1[:], arr2[:]...)
	fmt.Println("The merged array is: ", mergedArr)

}
