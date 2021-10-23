package main

import "reflect"

type  children struct {
	Age int
}
type Nested struct {
	X int
	Child children
}

func main(){
	vs := reflect.ValueOf(&Nested{}).Elem()
	vz := vs.Field(1)
	vz.Set(reflect.ValueOf(children{ Age:19 }))
}