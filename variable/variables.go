package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intnum int64 = 32767
	intnum = intnum + 1
	fmt.Println(intnum)
	var mystring string = "2456789"

	fmt.Println("this is the string part ")
	fmt.Println(` does this work well 
		idk man gotta execure and check `)

	fmt.Println(len(mystring))
	fmt.Println(utf8.RuneCountInString(mystring))

	var boolean1 bool = false

	fmt.Println(boolean1)

	otha := 1
	otha2 := 4
	fmt.Println(otha, otha2)
}
