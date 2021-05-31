package main

import (
	"fmt"
	"go_tutorial/nomadCoder/2.dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	// 1. Search Test
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// 2. Add Test
	word := "hello"
	definition2 := "Greeting"
	err2 := dictionary.Add(word, definition2)
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)

	err3 := dictionary.Add(word, definition2)
	if err3 != nil {
		fmt.Println(err3)
	}
	hello2, _ := dictionary.Search(word)
	fmt.Println(hello2)

	// 3. Update Test
	errUpdate := dictionary.Update(word, "New")
	if errUpdate != nil {
		fmt.Println(errUpdate)
	}
	newWord, _ := dictionary.Search(word)
	fmt.Println(newWord)

	// 4. Delete Test
	errDelete := dictionary.Delete(word)
	if errDelete != nil {
		fmt.Println(errDelete)
	}
	temp, err4 := dictionary.Search(word)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(temp)
	}

	fmt.Println(dictionary)
}
