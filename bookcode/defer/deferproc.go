package main

import "fmt"

func add(a, b int) {
	fmt.Println("add:" , a + b)
}

func f() {
	for i:=0;i<1;i++{
		defer add(1, 2)
	}
}
func main() {
	f()
}