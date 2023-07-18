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

// node represents a data node for our linked list.
// b 	byte stored
// p 	pointer to next data in linked list
type node struct {
	b byte
	p *node
}

var list *node

func init() {
	// 'ln' is the last node before 'nn' added
	var ln *node

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 'nn' is a new node
			var nn node
			// first node in the list
			if list == nil {
				list = &nn
			}
			// if there's a last node in the linked list,
			// have that last node point to the nn just added
			if ln != nil {
				ln.p = &nn
			}
			// make the nn the ln
			ln = &nn
			// Add a value to all even elements.
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				nn.b = 0xFF
			}
		}
	}

	// Count the number of elements in the link list.
	var ctr int
	n := list
	for n != nil {
		ctr++
		n = n.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

func main() {
	fmt.Println(LinkedListTraverse())
	fmt.Println(ColumnTraverse())
	fmt.Println(RowTraverse())

}

// LinkedListTraverse traverses the linked list linearly.
func LinkedListTraverse() int {
	var ctr int

	n := list
	for n != nil {
		if n.b == 0xFF {
			ctr++
		}

		n = n.p
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
