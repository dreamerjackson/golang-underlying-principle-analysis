package main

import "fmt"

func main(){
	defer catch("main")
	a()
}

func a() {
	defer b()
	panic("a panic")
}

func b() {
	defer fb()
	panic("b panic")
}

func fb() {
	defer catch("fb")
	panic("fb panic")
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}

