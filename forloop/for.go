package main

import "fmt"

func main() {
	mapmap2 := map[string]uint8{"adnan": 21, "chill": 22}

	intarr := []int32{1, 2, 3}

	for name, age := range mapmap2 {
		fmt.Printf("Name %v and age %v \n", name, age)
	}

	for index, value := range intarr {
		fmt.Printf("Index %v and value %v \n", index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
