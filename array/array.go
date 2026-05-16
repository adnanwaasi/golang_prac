package main

import "fmt"

func main() {
	var intarr [3]int32

	fmt.Println(intarr[0])
	fmt.Println(intarr[1:3])

	ass := []int32{1, 2, 3}
	fmt.Println(ass)
	fmt.Printf("the legnth %v and capacity %v", len(ass), cap(ass))
	ass = append(ass, 4)
	fmt.Println(ass)
	fmt.Printf("the legnth %v and capacity %v", len(ass), cap(ass))
}
