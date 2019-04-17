package main

import (
	"fmt"
	"log"
	"os"
)

type Vertex struct {
	X, Y int
}

func main() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{6, 8}
	)

	var console = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	fmt.Println("Recopie de v2 dans v3")
	v3 := v2 // Recopie de valeur
	fmt.Println("v1:", v1, "v2:", v2)
	fmt.Println("v3:", v3, "v2:", v2)
	fmt.Println("Changement de valeur pour v2.X")
	v2.X = 0
	fmt.Println("v3:", v3, "v2:", v2)

	fmt.Println("Redéfinition de v1")
	v1 = Vertex{3, 5} // Redéfinition de v1

	fmt.Printf("Pointeur de v1: %p, valeur de v1: %v\n", &v1, v1)

	// C'est fixe comme en C
	tableau := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(tableau)

	fmt.Println("On joue avec les slices")
	// C'est une slice, un tableau dynamique
	var slice []int // Il est vide
	printSlice(console, slice)
	slice = append(slice, 1) // On lui ajoute une valeur
	printSlice(console, slice)
	slice = append(slice, 2, 3) // On lui ajoute une valeur
	printSlice(console, slice)
	slice = slice[:0] // Nettoyage de la slice
	printSlice(console, slice)

	// Une slice avec make
	fmt.Println("Test de la fonction make")
	var slice2 = make([]int, 5)
	printSlice(console, slice2)

	// Instruction range
	for _, v := range tableau {
		fmt.Println(v)
	}

	// On fait une map (hash table)
	var m = map[string]Vertex{
		"Bell Labs": {40, 48},
		"Dell Inc.": {2, 4},
	}

	fmt.Println(m)
	delete(m, "Dell Inc.")
	fmt.Println(m)
	delete(m, "Dell")
	fmt.Println(m)
}

func printSlice(l *log.Logger, s []int) {
	l.Printf("len=%d, cappacity=%d %v\n", len(s), cap(s), s)
}
