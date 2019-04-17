package main

import "fmt"

func main() {
	// c'est une LIFO
	defer fmt.Println("world 1")
	defer fmt.Println("world 2")
	defer fmt.Println("world 3")

	fmt.Println("hello")
}
