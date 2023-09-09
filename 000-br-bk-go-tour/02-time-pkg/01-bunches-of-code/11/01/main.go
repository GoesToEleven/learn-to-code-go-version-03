package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	t := time.Now()
	fmt.Println(t)

	for {
		fmt.Print("> ")
		if !s.Scan() {
			break
		}
		fmt.Println(t.Format(s.Text()))
	}
}

/*
Reading from Standard Input
The if !scn.Scan() line is a crucial part of the program.
The Scan method scans a line from the input (os.Stdin in this case) up to a new line (\n) or until an error occurs.

If scanning is successful, Scan returns true.
If it encounters an error or the end of the file (EOF), it returns false.

The ! operator negates the Boolean value.
So !scn.Scan() will be true when Scan returns false,
which would indicate that the program should break out of the loop.
*/
