package main

import "fmt"

var g = 100

func f() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g
}

func main() {
	i := f()
	fmt.Printf("main: i = %d, g = %d\n", i, g)
}