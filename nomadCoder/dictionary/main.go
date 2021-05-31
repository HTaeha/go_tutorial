package main

import (
	"fmt"
	"go_tutorial/nomadCoder/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	// Search Test
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
