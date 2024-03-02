package main

import "fmt"

func main() {
	const name string = "World"
	var exclamationMark string = "!"
	var exclamationMark2 = "!"
	exclamationMark3 := "!"
	fmt.Println(hello() + " " + name + exclamationMark + exclamationMark2 + exclamationMark3)
}

func hello() string {
	return "Hello"
}
