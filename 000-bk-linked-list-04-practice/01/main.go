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
	var liList linkedList
	fmt.Printf("%T \t %v \n", liList, liList)

	var ll *node
	// ll := &node{}
	// ll := new(node)
	fmt.Printf("%T \t\t %v \n\n", ll, ll)

	liList.insertAtEnd(1)
	liList.insertAtEnd(2)
	liList.insertAtEnd(3)
	liList.printList()
}

func (list *linkedList) insertAtEnd(data int) {
	// 'nn' is a new node
	nn := &node{data, nil}
	if list.head == nil {
		list.head = nn
	} else {
		// iterate through list
		// 'in' is ITERATION NODE
		// loop will exit when we reach last node in list
		// last node in list will have 'nil' for 'next' - it won't point to the next node b/c it's the last
		in := list.head
		for in.next != nil {
			in = in.next
		}
		in.next = nn
	}
}

func (list *linkedList) printList() {
	n := list.head
	for n != nil {
		fmt.Printf("%p \t %d \t %v \n", n, n.data, n)
		n = n.next
	}
}
