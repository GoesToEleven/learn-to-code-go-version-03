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

	fmt.Print("\nTHIS WORKS \n TYPE T \n RECEIVER T\n")
	interfaceimplementation.ThisWorks()

	fmt.Print("\nTHIS DOES NOT WORK \n TYPE T \n RECEIVER *T\n")
	fmt.Print("The interface is not implemented\n")
	interfaceimplementation.ThisDoesNotWork()

	fmt.Print("\nTHIS WORKS \n TYPE *T \n RECEIVER T and RECEIVE *T\n")
	interfaceimplementation.ThisAlsoAllWorks()
}