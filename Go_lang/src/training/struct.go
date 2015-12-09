package main

import (
	"fmt"
	"log"
)

type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return p.name
}
func (p person) getAge() int {
	return p.age
}

func main() {
	dictionary := map[string]person{
		"0": {"Jai", 10},
		"1": {"Jai2", 210},
	}

	for _, p1 := range dictionary {
		fmt.Println(p1.getAge())
		log.Println("")
	}

}
