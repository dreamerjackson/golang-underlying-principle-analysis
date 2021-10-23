package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("jonson ReflectCallFunc")
}

func main(){
	user := User{1, "jonson", 25}
	getType := reflect.TypeOf(user)
	fmt.Println("get Type is :", getType.Name())
	getValue := reflect.ValueOf(user)
	fmt.Println("get all Fields is:", getValue)
}


