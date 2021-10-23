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

func (u User) RefCallArgs( age int, name string) error {
	return nil
}

func (u User) RefCallNoArgs() {
	fmt.Println("hello world")
}

func (u User) RefCallPoint(name string, age *int) {
	fmt.Println(" name: ", name, ", age:", *age)
}

func (u *User) RefPointMethod() {
	fmt.Println("hello world")
}

func (u *User) PointMethodReturn(name string, age int)(string, int) {
	return name,age
}


func main(){
	user := User{1, "jonson", 25}
	ref := reflect.ValueOf(user)
	tf:= ref.MethodByName("RefCallArgs").Type()
	fmt.Printf("numIn:%d,numOut:%d,numMethod:%d\n",tf.NumIn(), tf.NumOut(),ref.NumMethod())


	m := ref.MethodByName("RefCallNoArgs")
	args := make([]reflect.Value, 0)
	m.Call(args)

	m = ref.MethodByName("RefCallArgs")
	args  = []reflect.Value{reflect.ValueOf(18),reflect.ValueOf("json")}
	m.Call(args)



	m = ref.MethodByName("RefCallPoint")
	age := 19
	args = []reflect.Value{reflect.ValueOf("jonson"), reflect.ValueOf(&age)}
	m.Call(args)

	refnew := reflect.ValueOf(&user)
	m = refnew.MethodByName("RefPointMethod")
	args  = make([]reflect.Value, 0)
	m.Call(args)

	m  = refnew.MethodByName("PointMethodReturn")
	args = []reflect.Value{reflect.ValueOf("jonson"), reflect.ValueOf(30)}
	res:= m.Call(args)
	fmt.Println("return name:",res[0].Interface())
	fmt.Println("return age:",res[1].Interface())
}