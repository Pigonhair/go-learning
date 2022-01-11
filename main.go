package main

import (
	"fmt"

	"github.com/Pigonhair/go-learning/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	dictionary.Search(baseWord)
	fmt.Println(dictionary)
	dictionary.Delete(baseWord)
	word, _ := dictionary.Search(baseWord)
	fmt.Println(dictionary)
	fmt.Println(word)
}
