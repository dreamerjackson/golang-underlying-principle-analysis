package main

import (
	"os"
	"runtime/pprof"
)

func main(){
	a(99)
}

func a(num int){
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}

