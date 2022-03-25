package modificationModule

import (
	"fmt"
	"go-study/person"
)

func Modify(person person.Person) {
	person.Age = 100
	person.Name = "Justin Bieber"

	fmt.Println(person)
}
