package main

func small() string {
	s := "hello, " + "world!"
	return s
}

func fib(index int) int{

	if index <2{
		return index
	}

	return fib(index-1) + fib(index-2)
}

func main() {
	small()
	fib(65)
}

/*
	» go tool compile -m=2 5_inline.go                                                                     jackson@jacksondeMacBook-Pro
	5_inline.go:3:6: can inline small as: func() string { s := "hello, world!"; return s }
	5_inline.go:8:6: cannot inline fib: recursive

*/