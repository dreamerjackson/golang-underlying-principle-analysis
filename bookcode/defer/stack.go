package main

import "fmt"

func add(a, b int) {
	fmt.Println("add:" , a + b)
}

func f() {
		defer add(3, 4)
		add(3,4)
}
func main() {
	f()
}