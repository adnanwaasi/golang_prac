package main

import "fmt"

func main() {
	var mapmap map[string]uint8 = make(map[string]uint8)
	fmt.Println(mapmap)
	mapmap2 := map[string]uint8{"waasi": 21, "chill": 8}
	fmt.Println(mapmap2)
	fmt.Println(mapmap2["chill"])
	delete(mapmap2, "chill")
	fmt.Println(mapmap2)

	for name, age := range mapmap2 {
		fmt.Printf("Index: %v value: %v \n", name, age)
	}
}
