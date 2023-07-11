package main

import "fmt"

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

func (ll *sll) Display() {
	if ll.head == nil {
		fmt.Println("Sorry there is no element in this linked list")
	} else {
		current := ll.head
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func (ll *sll) addNext(data int, val int) {
	if ll.head == nil {
		ll.head = &node{data, nil}
	} else {
		current := ll.head

		for current.next != nil && current.data != val {
			current = current.next
		}
		temp := current.next
		current.next = &node{data, temp}
	}
}

func (ll *sll) addPrev(data int, val int) {
	if ll.head == nil {
		ll.head = &node{data, nil}
	} else {
		current := ll.head

		for current.next != nil && current.next.data != val {
			current = current.next
		}
		temp := current.next
		current.next = &node{data, temp}
	}
}

func main() {
	var ll sll
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Display()
	ll.addNext(6, 2)
	ll.Display()
	ll.addPrev(71, 6)
	ll.Display()
}
