package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func main() {
	var ll linkedList
	fmt.Printf("%v \n", ll)

	ll.insertAtEnd(1)
	ll.insertAtEnd(2)
	ll.insertAtEnd(3)
	ll.printList()
}

func (list *linkedList) insertAtEnd(data int) {
	nn := &node{data, nil}
	if list.head == nil {
		list.head = nn
	} else {
		// find the last node in the list
		cn := list.head
		for cn.next != nil {
			cn = cn.next
		}
		cn.next = nn
	}
}

func (list *linkedList) printList() {
	// start at the beginning
	cn := list.head
	for cn != nil {
		fmt.Printf("this nodes address: %p \t data: %d \t next node's address: %p \n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
