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

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
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

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(p, data)
		if i > 0 {
			if _, err := store(s, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	e := &Eagle{
		Host:    "localhost:8000",
		Timeout: time.Second,
	}
	o := &Optima{
		Host:    "localhost:9000",
		Timeout: time.Second,
	}

	if err := Copy(e, o, 3); err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("ran")
}
