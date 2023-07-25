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

	// this is deprecated:
	// rand.Seed(time.Now().UnixNano())
}

// Data we are copying.
type Data struct {
	Line string
}

// Eagle is the system we pull Data from.
type Eagle struct {
	Host    string
	Timeout time.Duration
}

// Optima is the system we store Data in.
type Optima struct {
	Host    string
	Timeout time.Duration
}

// System wraps Eagle and Optima together.
type System struct {
	Eagle
	Optima
}

// Pull pulls Data from Eagle.
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

// pull pulls Data from Eagle.
func pull(e *Eagle, data []Data) (int, error) {
	for i := range data {
		if err := e.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Store stores Data in Optima.
func (*Optima) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// store stores Data in Optima.
func store(o *Optima, data []Data) (int, error) {
	for i := range data {
		if err := o.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// Copy pulls and stores Data from System.
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(&sys.Eagle, data)
		if i > 0 {
			if _, err := store(&sys.Optima, data[:i]); err != nil {
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
