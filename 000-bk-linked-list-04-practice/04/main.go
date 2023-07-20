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
	var mylist linkedList
	mylist.addNodeToEnd(1)
	mylist.addNodeToEnd(2)
	mylist.addNodeToEnd(3)
	mylist.addNodeToEnd(4)

	mylist.printList()
}

func (list *linkedList) addNodeToEnd(data int) {
	nn := &node{data, nil}
	if list.head == nil {
		list.head = nn
	} else {
		// find end of the list
		// start at the beginning
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
		fmt.Printf("current node:%p \t data: %d \t next node: %p \n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
