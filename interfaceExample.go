package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	//creating the type shape to allow access to the area method
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
	//method attached to rectangle that calculates the area
	//implements the area interface
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	//passing circle and rectangle into this as an argument
	//automatically casts it as an implementation of the shape type
	return s.area()
}

func Shape() {
	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape{c1, r1}
	for _, shape := range shapes {
		// _ represents the would be index
		// shape is the value, ignoring the index?
		fmt.Println(getArea(shape))
	}
}
