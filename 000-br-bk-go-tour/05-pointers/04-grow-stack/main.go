package main

const size = 1

func main() {
	s := "HELLO"
	stackCopy(0, &s, [size]int{})
}

//go:noinline
func stackCopy(c int, s *string, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(c, s, a)
}
