package main

import (
	"fmt"
	"mymodule/000-ae-method-sets/interfaceimplementation"
	"mymodule/000-ae-method-sets/thisallworks"
)

/*
The method set
of a defined type T
methods declared with
receiver type T.

The method set
of a pointer to a defined type T
(where T is neither a pointer nor an interface)
methods declared with
receiver *T or T.
*/

func main() {
	thisallworks.RunMe()

	fmt.Printf("\nTHIS WORKS \n TYPE T \n RECEIVER T")
	interfaceimplementation.ThisWorks()

	fmt.Printf("\nTHIS DOES NOT WORKS \n TYPE T \n RECEIVER *T")
	fmt.Printf("\nThe interface is not implemented")
	interfaceimplementation.ThisDoesNotWork()

	fmt.Printf("\nTHIS WORKS \n TYPE *T \n RECEIVER T and RECEIVE *T")
	interfaceimplementation.ThisAllWorks()
}
