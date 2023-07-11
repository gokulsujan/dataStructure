package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type sll struct {
	head *node
}

func (ll *sll) Add(data int) {
	newNode := node{data, nil}
	if ll.head == nil {
		ll.head = &newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
}

func arrayToLl(arr []int, ll *sll) {
	for i := 0; i < len(arr); i++ {
		ll.Add(arr[i])
	}

}

func (ll *sll) Display() {
	if ll.head == nil {
		fmt.Println("This linked list is empty")
	} else {
		current := ll.head
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ll sll
	arrayToLl(arr, &ll)
	ll.Display()
}
