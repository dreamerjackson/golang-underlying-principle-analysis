package main

import "fmt"

var g = 100

func f() (r int) {
	r = g
	defer func() {
		r = 200
	}()

	fmt.Printf("f: r = %d\n", r)

	r = 0
	return r
}

func main() {
	i := f()
	fmt.Printf("main: i = %d, g = %d\n", i, g)
}