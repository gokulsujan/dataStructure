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

	fmt.Println("How you want to delete: ")
	fmt.Println("Press 1 to delete using array index")
	fmt.Println("Press 2 to delete using array value")
	var delOpt int
	fmt.Scan(&delOpt)
	if delOpt == 1 {
		fmt.Println("Enter the array index you want to delete: ")
		var index int
		fmt.Scan(&index)
		if index >= limit {
			fmt.Println("Sorry you entered the index value greater than the array size")
		} else {
			for i := index; i < limit; i++ {
				if i == limit-1 {
					arr[i] = 0
				} else {
					arr[i] = arr[i+1]
				}
			}
		}
	} else if delOpt == 2 {
		fmt.Println("Enter the array value you want to delete: ")
		var value int
		fmt.Scan(&value)
		for i := 0; i < limit; i++ {
			if arr[i] == value {
				for j := i; j < limit; j++ {
					if j == limit-1 {
						arr[j] = 0
					} else {
						arr[j] = arr[j+1]
					}
				}
			} else {
				continue
			}
		}
	} else {
		fmt.Println("Invalid option, operation cancelled")
	}
	fmt.Println(arr)

}
