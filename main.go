package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// name := "nico"
	fmt.Println("HELLO WORLD")
	fmt.Println(multiply(2, 2))
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)
}
