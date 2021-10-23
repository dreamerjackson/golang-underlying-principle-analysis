package main

func add(a int,b int) int{
	return a*b
}

func mul(a int,b int) int{
	return add(a,b)
}

func main(){
	mul(3,4)
}