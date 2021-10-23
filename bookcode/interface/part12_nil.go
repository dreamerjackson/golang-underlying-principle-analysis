package main

import (
	"fmt"
	"os"
)


func foo() error {
	var err error // nil
	return err
}

func foo_path() error {
	var err *os.PathError // nil
	return err
}

func main(){
	err := foo()
	if err == nil {
		fmt.Println("foo() err == nil true")
	}
	err = foo_path()
	if err == nil {
		fmt.Println("foo_path() err == nil true")
	}
	err = wrapDo()
	if err == nil {
		fmt.Println("wrapDo err == nil true")
	}
}

func do() *os.PathError {
	return nil
}
func wrapDo() error {
	return do()
}
