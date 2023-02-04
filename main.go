package main

import (
	"fmt"

	"github.com/Esppe0w/goproject_1/dict"
)

func main() {
	dictionary := dict.Dictionary{}

	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

}
