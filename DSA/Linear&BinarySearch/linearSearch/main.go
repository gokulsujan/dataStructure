package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 2

	index := linearSearch(arr, target)
	fmt.Println(index)
}
