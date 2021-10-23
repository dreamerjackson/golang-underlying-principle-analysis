package main

var count = 0
func add() {
	count++
}

func main() {
	go add()
	go add()
}