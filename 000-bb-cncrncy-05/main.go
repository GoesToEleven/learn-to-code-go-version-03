package main

func main() {
	var ch chan int
	<-ch
	ch <- 3
}
