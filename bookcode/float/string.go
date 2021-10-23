package main

import "fmt"

var arr [3]int
var arr2  = [4]int{1,2,3,4}
var arr3  =[...]int{2,3,4}
func main(){
	fmt.Printf("类型arr2: %T,类型arr3: %T\n",arr2,arr3,arr[99])


}