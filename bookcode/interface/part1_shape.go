package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
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

type Triangle struct {
	a, b, c float64
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (r Rectangle) getHeight() float64 {
	return  r.a
}

func main(){
	var s Shape
	s = Rectangle{3, 4}
	fmt.Printf("长方形周长:%v, 面积:%v \n",s.perimeter(),s.area())
	s = Triangle{3, 4, 5}
	fmt.Printf("三角形周长:%v, 面积:%v",s.perimeter(),s.area())
}