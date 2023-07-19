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
	ll := linkedList{}
	fmt.Printf("%v \t %T \n", ll, ll)

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
		cn := list.head
		for cn.next != nil {
			cn = cn.next
		}
		cn.next = nn
	}
}

func (list *linkedList) printList() {
	cn := list.head
	for cn != nil {
		fmt.Printf("%p \t %d \t %p \n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
