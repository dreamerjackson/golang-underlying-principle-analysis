package main

import (
	"fmt"
	"sync"
)

func main(){
	for i:=1;i<5;i++{
		defer fmt.Println("start ",i)
	}
	var m sync.Mutex
	var n sync.Mutex
	m.Lock()
	defer m.Lock()
	n.Lock()
	defer n.Lock()
}

