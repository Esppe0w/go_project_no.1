package main

import (
	"fmt"

	"github.com/Esppe0w/goproject_1/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "Frist word"}
	difinition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(difinition)

}
