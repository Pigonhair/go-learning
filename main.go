package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") //defer은 func이 값을 return시킨 다음에 실행됨
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	// totalLength, up := lenAndUpper("kunwoo")
	// fmt.Println(totalLength, up)
	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)
	// result2 := canIDrink(16)
	// fmt.Println(result2)
	a := 2
	b := &a //메모리주소를 참조함 (a의 pointer)
	*b = 20 //b를 가지고 a의 값을 바꿈
	fmt.Println(a)
	// *b = 202020
	// fmt.Println(a)
}
