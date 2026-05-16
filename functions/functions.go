package main

import ("fmt"
	errors")

func main() {
	fmt.Printf("the ans is num %v and %v", intdiv(87, 6))
}

func intdiv(num int, den int) (int, int) {
	return (num / den), (num % den)
}
