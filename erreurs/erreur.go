package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("42")

	if err != nil {
		log.Panicln("i is nil")
	}

	fmt.Printf("sucess: %v", i)

	n := math.Sqrt(-1)
	if math.IsNaN(n) {
		log.Panicf("cannot compute sqrt(%s)", n)
	}
	fmt.Println(n)
}
