package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

var (
	ErrFour = New("my fourth error")
	ErrFive = New("my fifth error")
)

func anotherCall() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(2) == 0 {
		return ErrFour
	} else {
		return ErrFive
	}
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

	if err := anotherCall(); err != nil {
		switch err {
		case ErrFour:
			fmt.Println(err)
			return
		case ErrFive:
			fmt.Println(err)
			return
		}
	}

}
