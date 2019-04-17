package main

import (
	"fmt"
	"log"
	"reflect"

	"m2i.fr/figgeo/calcmath"
	"m2i.fr/figgeo/cercle"
	"m2i.fr/figgeo/rectangle"
)

/*
 * Fonction d'initialisation exécutée avant main()
 */
func init() {
	fmt.Println("Actions de pré-lancement, avant le main")
}

/*
 * Fonction principale de l'application
 */
func main() {
	r, err := rectangle.New(2, 34)
	c := cercle.Cercle{Rayon: 2.0}

	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(*r)
	var new_value uint = 4
	r.Resize(&new_value, &new_value)
	fmt.Println(*r)
	r.Resize(nil, &new_value)
	fmt.Println(*r)

	PrintSurface(r)
	PrintSurface(c)
}

func PrintSurface(f calcmath.CalcSurface) {
	fmt.Printf("Surface calculée de %v %v\n", reflect.TypeOf(f), f.GetSurface())
}
