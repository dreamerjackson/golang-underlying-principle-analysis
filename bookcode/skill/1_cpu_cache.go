package main

import (
	"fmt"
	"time"
)

func main() {
	var arr [256][256]int
	var arr2 [256][256]int

	before1 := time.Now()
	for i:=0;i<256;i++{
		for j:=0;j<256;j++{
			arr[i][j] = i*j
		}
	}
	after1:= time.Now()
	fmt.Println(after1.Sub(before1))

	before2 := time.Now()
	for i:=0;i<256;i++{
		for j:=0;j<256;j++{
			arr2[j][i] = i*j
		}
	}
	after2:= time.Now()
	fmt.Println(after2.Sub(before2))


}
