package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // target not found
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 2

	index := binarySearch(arr, target)
	fmt.Println(index)
}
