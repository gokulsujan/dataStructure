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

func (ll *sll) addBegining(data int) {
	ll.head = &node{data, ll.head}
}

func main() {
	var ll sll
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Display()
	ll.addBegining(25)
	ll.Display()

}
