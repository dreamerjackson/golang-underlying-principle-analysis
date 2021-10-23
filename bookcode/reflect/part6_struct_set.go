package main

import "reflect"

var s struct {
	X int // an exported field
	y float64 // a non-exported field
}


func main(){
	vs := reflect.ValueOf(&s).Elem()
	vx:= vs.Field(0)
	vb := reflect.ValueOf(123)
	vx.Set(vb)
}