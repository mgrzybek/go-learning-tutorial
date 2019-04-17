package main

import (
	"fmt"
	"log"
	"math"
)

/***************************************************************/
// La classe
type Vertex struct {
	X, Y float64
}

// La méthode est une fonction qui a comme référence la classe
// Méthode Abs() de la classe Vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/***************************************************************/
type entier int

// Méthode Abs() du type entier
func (e entier) Abs() uint {
	if e >= 0 {
		log.Println("On ne change pas e")
		return uint(e)
	}

	log.Println("On change e")
	return uint(-e)
}

/***************************************************************/
type Abser interface {
	Abs() float64
}

/***************************************************************/
func main() {
	var e entier = -3
	var i interface{} = "hello"
	var j interface{} = 1

	v := Vertex{3, 4}
	s := i.(string)
	r := j.(int)

	fmt.Println(e.Abs())
	fmt.Println(v.Abs())

	fmt.Println(s)
	fmt.Println(r)
}
