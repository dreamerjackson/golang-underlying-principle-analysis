package main

import (
	"fmt"
	"math"
)

func main() {
	var z float32
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

func NaN2(){
	var number float64 = math.NaN()
	bits := math.Float64bits(number)

	binary := fmt.Sprintf("%.64b", bits)

	fmt.Println("Bit Pattern:",binary)
}
