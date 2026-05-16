package main

import (
	"fmt"
	"strings"
)

func main() {
	var fahh string = "résumé"
	hahh := fahh[0]
	fmt.Printf("%v,%T\n", hahh, hahh)
	for i, v := range fahh {
		fmt.Println(i, v)
	}

	var builder strings.Builder
	strslice := []string{"m", "a", "x"}
	for i := range strslice {
		builder.WriteString(strslice[i])
	}

	fi := builder.String()

	fmt.Println("\n ", fi)
}
