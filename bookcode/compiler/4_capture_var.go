package main

import "fmt"

func main(){
	a := 1
	b := 2
	go func() {
		fmt.Println(a,b)
	}()
	a = 99
}
/*
 » go tool compile -m=2  main.go | grep capturing
main.go:9:15: main.func1 capturing by ref: a (addr=true assign=true width=8)
main.go:9:17: main.func1 capturing by value: b (addr=false assign=false width=8)
*/