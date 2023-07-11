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
	newNode := &node{data, nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *sll) Display() {
	if ll.head == nil {
		fmt.Println("The ll is empty")
	} else {
		current := ll.head
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func (ll *sll) Delete(val int) {
	current := ll.head.next

	if ll.head.data == val {
		ll.head = current
	} else {
		for current != nil {
			if current.next.data == val {
				del := current.next
				current.next = del.next
				break
			} else {
				current = current.next
				continue
			}
		}
	}
}

func main() {
	var ll sll
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Display()
	fmt.Println("Enter the value to delete")
	var val int
	fmt.Scan(&val)
	ll.Delete(val)
	ll.Display()
}
