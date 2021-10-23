package main
// go run  -gcflags="-l" stack_trace.go
func trace(arr []int,a int,b int) int{
	panic("test trace")
	return 0
}

func main(){
	arr := []int{1,2,3}
	trace(arr,5,6)
}