/*
https://tour.golang.org/welcome/1
_test_
@author Mathieu GRZYBEK
*/
package main

import (
	"fmt"
)

func main() {
	var a int = 2
	var b int = 3
	var c, python, java = true, false, "no!"
	complex := 1 + 2i

	const toto = 0

	fmt.Println("Hello, 世界")
	fmt.Println("add:", add(1, 2))
	fmt.Println("sub:", sub(1, 2))
	fmt.Println(a, b)
	fmt.Println("a:", &a, "b:", &b)

	fmt.Println(split(100))

	fmt.Println(a, b, c, python, java)
	fmt.Println("a:", uint(a))
	fmt.Println(complex)

	Big := 1 << 10
	fmt.Println("Décalage de bits:", Big)

	Small := 1 >> 10
	fmt.Println("Décalage de bits:", Small)

	for i := 0; i < 10; i++ {
		fmt.Println("compteur:", i)
	}
}

func swap(x, y string) (string, string) { return y, x }

func add(x int, y int) int { return x + y }
func sub(x, y int) int     { return x - y }

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
