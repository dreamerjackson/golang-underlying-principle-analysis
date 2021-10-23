package main

import "runtime/debug"

func main(){
	a(99)
}

func a(num int){
	debug.PrintStack()
}