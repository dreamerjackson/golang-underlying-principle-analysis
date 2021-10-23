package main

import (
	"fmt"
	"reflect"
)

func main(){
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	//

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)//type:int64 value:56
	b := "json"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)//type:string value:Naveen
	c := 12.5
	z := reflect.ValueOf(c).Float()
	fmt.Printf("type:%T value:%v\n", z, z)//type:string value:Naveen
}