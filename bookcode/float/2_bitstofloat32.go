package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = 0.085
	fmt.Printf("Starting Number: %f\n", number)
	// Float32bits returns the IEEE 754 binary representation
	bits := math.Float32bits(number)

	binary := fmt.Sprintf("%.32b", bits)
	bias := 127
	sign := bits & (1 << 31)
	exponentRaw := int(bits >> 23)
	exponent := exponentRaw - bias

	var mantissa float64
	for index, bit := range binary[9:32] {
		if bit == 49 {
			position := index + 1
			bitValue := math.Pow(2, float64(position))
			fractional := 1 / bitValue
			mantissa = mantissa + fractional
		}
	}

	value := (1 + mantissa) * math.Pow(2, float64(exponent))

	fmt.Printf("Sign: %d Exponent: %d (%d) Mantissa: %f Value: %f\n",
		sign,
		exponentRaw,
		exponent,
		mantissa,
		value)
}

