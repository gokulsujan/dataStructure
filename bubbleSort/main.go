package main

import "fmt"

func bubbleSort(arr []int) {
	limit := len(arr)

	for i := 0; i < limit-1; i++ {
		swap := false
		for j := 0; j < limit-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swap = true
			}
		}

		if swap == false {
			break
		}
	}
}

func main() {
	arr := []int{10, 20, 40, 30, 50}
	bubbleSort(arr)
	fmt.Println(arr)

}
