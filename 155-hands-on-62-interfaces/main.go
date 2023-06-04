package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{
		radius: 4,
	}

	s1 := square{
		length: 5,
		width:  8,
	}

	fmt.Println(info(c1))
	fmt.Println(info(s1))
}

/*
create a type SQUARE
length int
width int
create a type CIRCLE
radius float64
attach a method to each that calculates AREA and returns it
circle area= π r 2
math.Pi
math.Pow
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle*/
