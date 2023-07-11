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

	fmt.Println("How you want to search: ")
	fmt.Println("Press 1 to search using array index")
	fmt.Println("Press 2 to search using array value")
	var serOpt int
	fmt.Scan(&serOpt)

	if serOpt == 1 {
		fmt.Println("Enter the array index you want to search")
		var index int
		fmt.Scan(&index)
		if index >= limit {
			fmt.Println("Sorry you entered the index value greater than the array size")
		} else {
			fmt.Println("The value in the array index", index, "is ", arr[index])
		}
	} else if serOpt == 2 {
		fmt.Println("Enter the array value you want to search")
		var val int
		fmt.Scan(&val)
		found := 0
		for i := 0; i < limit; i++ {
			if arr[i] == val {
				fmt.Println("The array value ", val, "found at the index", i)
				found = 1
				break
			} else {
				found = 0
			}
		}
		if found == 0 {
			fmt.Println("Sorry the entered element is  not in the array")
		}

	} else {
		fmt.Println("Sorry invalid option")
	}
}
