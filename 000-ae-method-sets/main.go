package main

import (
	"fmt"
	"mymodule/000-ae-method-sets/interfaceimp"
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
	fmt.Print("\n Just methods\n")
	thisallworks.RunMe()

	fmt.Print("\n THIS WORKS \n TYPE T \n RECEIVER T\n")
	interfaceimplementation.ThisWorks()

	fmt.Print("\n THIS DOES NOT WORK \n TYPE T \n RECEIVER *T\n")
	interfaceimplementation.ThisDoesNotWork()

	fmt.Print("\n THIS WORKS \n TYPE *T \n RECEIVER T and RECEIVER *T\n")
	interfaceimplementation.ThisAlsoAllWorks()

	fmt.Println("interface imp")
	interfaceimp.Runnit()
}
