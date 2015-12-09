package main

import "fmt"

func main() {
	name := "HelloWorld"
	slc := []byte(name)
	for index, char := range slc {
		fmt.Printf("Index := %d Character := %c\n", index, char)
	}
}
