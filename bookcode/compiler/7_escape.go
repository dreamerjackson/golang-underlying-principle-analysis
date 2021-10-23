package main

var o *int
func main(){
	l := new(int)
	*l = 42
	m := &l
	n := &m
	o = **n
}