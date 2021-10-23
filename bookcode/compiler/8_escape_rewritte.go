package main

import "fmt"

func do() {
	a := 1
	func() {
		fmt.Println(a)
		a = 2
	}()
}


/*

func do() {
	a := 1
	func1(&a)
}
func func1(a *int) {
	fmt.Println(*a)
	*a = 2
}
*/