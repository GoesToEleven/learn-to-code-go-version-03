package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	t := time.Now()
	fmt.Println(t)

	for {
		fmt.Printf("> ")
		if !scn.Scan() {
			break
		}
		fmt.Println(t.Format(scn.Text()))
	}
}

// go doc time.Layout | more
// 9s and 0s also for precision in fractions of a second, up to 9 digitis
// padding with space "_2"
// 01/02 03:04:05 2006 -0700

/*
Reading from Standard Input
The if !scn.Scan() line is a crucial part of the program. The Scan method scans a line from the input (os.Stdin in this case) up to a new line (\n) or until an error occurs.

If scanning is successful, Scan returns true.
If it encounters an error or the end of the file (EOF), it returns false.
The ! operator negates the Boolean value. So !scn.Scan() will be true when Scan returns false, which would indicate that the program should break out of the loop.
*/
