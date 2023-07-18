package main

import "fmt"

// Create a square matrix of 16,777,216 bytes.
const (
	rows = 4 * 1024
	cols = 4 * 1024
)

// matrix represents a matrix with a large number of
// columns per row.
var matrix [rows][cols]byte

// data represents a data node for our linked list.
// v 	value stored
// p 	pointer to next data in linked list
type data struct {
	v byte
	p *data
}

// list points to the HEAD of the list.
var list *data

func init() {
	// last points to the TAIL of the list
	var last *data

	// Create a link list with the same number of elements.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			// Create a new node and link it in.
			var d data
			if list == nil { // sets the head of the linked list
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			// Add a value to all even elements.
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	// Count the number of elements in the link list.
	var ctr int
	d := list
	for d != nil {
		ctr++
		d = d.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

// LinkedListTraverse traverses the linked list linearly.
func LinkedListTraverse() int {
	var ctr int

	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}

		d = d.p
	}

	return ctr
}

// ColumnTraverse traverses the matrix linearly down each column.
func ColumnTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// RowTraverse traverses the matrix linearly down each row.
func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
