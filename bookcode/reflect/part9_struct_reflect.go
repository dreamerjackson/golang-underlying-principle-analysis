package main

import (
	"fmt"
	"reflect"
)

func MakeStruct(vals ...interface{}) reflect.Value {
	var sfs []reflect.StructField
	for k, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", (k + 1)),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so
}

func main(){

	sr := MakeStruct(0, "", []int{})

	fmt.Println(sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface())

	fmt.Println(sr.Elem().Field(1).Interface())
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println(sr.Elem().Field(1).Interface())

	fmt.Println(sr.Elem().Field(2).Interface())
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println(sr.Elem().Field(2).Interface())
}