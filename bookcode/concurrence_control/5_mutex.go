package main

import (
	"sync"
)

var count int64  = 0
var m sync.Mutex
func add() {
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	go add()
	go add()
}