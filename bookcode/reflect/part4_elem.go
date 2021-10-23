package main

import (
	"fmt"
	"reflect"
)

func main(){
	aa := 56
	xx := reflect.ValueOf(&aa).Elem().Int()
	fmt.Printf("type:%T value:%v\n", xx, xx) //type:int64 value:56

	//
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind()) // map array slice
	tx, ty := ta.Elem(), tb.Elem()
	// byte is an alias of uint8
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8

	//
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()
	fmt.Println("settability of pointer:", newValue.CanSet())
	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)
}



