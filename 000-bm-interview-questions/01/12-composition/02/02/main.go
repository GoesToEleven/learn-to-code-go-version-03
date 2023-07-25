package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Data struct {
	Line string
}

type Eagle struct {
	Host    string
	Timeout time.Duration
}

type Optima struct {
	Host    string
	Timeout time.Duration
}

// system embeds Eagle and Optima
// so it will implement both Puller & Storer interface
// and also the PullStorer interface
type System struct {
	Eagle
	Optima
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type PullStorer interface {
	Puller
	Storer
}

func (*Eagle) Pull(d *Data) error {
	switch r.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("error reading data from Eagle")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// change paramenter to PULLER
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func (*Optima) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// change paramenter to STORER
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Eagle: Eagle{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Optima: Optima{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("ran")
}
