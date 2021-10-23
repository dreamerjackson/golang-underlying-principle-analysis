package main

import (
	"fmt"
	"reflect"
)

func main(){
	var num float64 = 1.2345
	fmt.Println("type: ", reflect.TypeOf(num))           //type:  float64
	fmt.Println("value: ", reflect.ValueOf(num))         //value:  1.2345

	equl := reflect.TypeOf(num).Kind() == reflect.Float64
	fmt.Println("kind is float64: ", equl)  //kind is float64: true
}
