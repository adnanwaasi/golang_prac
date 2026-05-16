package main

import (
	"fmt"
	"math"
)

func main() {
	thing1 := [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("the memory location of the the thing1 is %p", &thing1)
	result := square(&thing1)
	fmt.Printf("\nthe result is %v", result)
	fmt.Printf("result address %p", &result)
	fmt.Println("\nthe value of thing1 is ", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n the memory address of thing2 array is %p", thing2)

	for i := range thing2 {
		thing2[i] = math.Pow(thing2[i], 2)
	}
	return *thing2
}
