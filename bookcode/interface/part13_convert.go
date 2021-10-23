package main

import (
	"fmt"
	"math"
)
type Shape interface {
	perimeter() float64
	area() float64
}

type Areaifer interface {
	area() float64
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) perimeter() float64 {
	return (r.a + r.b) * 2
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

func (r Rectangle) getHeight() float64 {
	return  r.a
}
func main(){
	var shape Shape
	var area Areaifer
	shape = Rectangle{3, 4}
	area = shape
	fmt.Println("area:",area.area())
}