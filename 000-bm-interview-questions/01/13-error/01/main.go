package main

import "fmt"

// pkg builtin
type error interface {
	Error() string
}

// pkg errors errors.go
type errorString struct {
	s string
}

// pkg errors errors.go
func (e *errorString) Error() string {
	return e.s
}

// pkg errors errors.go
func New(text string) error {
	return &errorString{text}
}

func someCall() error {
	return New("my third error")
}

func main() {
	var a error
	v := errorString{"my error"}
	fmt.Println(v.Error())

	a = &v
	fmt.Println(a.Error())

	b := New("my second error")
	fmt.Println(b.Error())

	if err := someCall(); err != nil {
		fmt.Println(err)
	}

}
