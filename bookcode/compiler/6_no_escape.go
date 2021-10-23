package main

func main() {
	f()
}

func f() int{
	var z *int
	a := 1
	z = &a
	return *z
}