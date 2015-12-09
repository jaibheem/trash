package main

import "fmt"

func main() {
	kvpair := map[string]string{
		"1": "One",
		"2": "Two",
		"3": "Three",
	}
	for key, value := range kvpair {
		fmt.Printf("Key = %s, Value = %s\n", key, value)
	}
	kvpair["4"] = "Four"
	fmt.Println("==================")
	for key, value := range kvpair {
		fmt.Printf("Key = %s, Value = %s\n", key, value)
	}
}
