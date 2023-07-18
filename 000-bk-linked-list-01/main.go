package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	list := LinkedList{}
	list.insertAtEnd(1)
	list.insertAtEnd(2)
	list.insertAtEnd(3)
	list.printList()
}

func (list *LinkedList) insertAtEnd(d int) {
	// nn is a new node
	nn := &Node{
		data: d,
		next: nil,
	}

	// if our linked list has no nodes, the nn is the linked list head
	if list.head == nil {
		list.head = nn

	} else {
		// loop through, from beginning to end, our linked list
		// the variable 'in' represents the ITERATION NODE in our list
		// when we start the iteration, 'in' is the first node in the list (list.head)
		in := list.head
		// for each iteration, if ITERATION NODE has an 'in.next' pointing to a next node (eg, the pointer is not nil)
		for in.next != nil {
			// set 'in' to the address of the next node in the list
			in = in.next
		}
		// when in.next == nil
		// the loop above will exit
		// this is the node that was the LAST NODE in the list
		// we now have a NEW NODE 'nn' however,
		// so for this last ITERATION NODE that currently does not point to a node beyond it
		// set it to point to the NEW NODE 'nn'
		in.next = nn

		// the NEW NODE 'nn' we created will now be the last node
		// and this NEW NODE will not have a pointer to a next node
		// as this NEW NODE is the new LAST NODE
	}
}

func (list *LinkedList) printList() {
	// n is a node
	n := list.head
	for n != nil {
		fmt.Printf("address: %p \tdata: %d \tstructure: %v\n", n, n.data, n)
		n = n.next
	}
}
