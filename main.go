package main

import (
	"fmt"
	"go-study/modificationModule"
	"go-study/person"
)

func main() {
	var person = person.Person{Name: "name", Age: 1}

	modificationModule.Modify(person)
	fmt.Println(person)
}
