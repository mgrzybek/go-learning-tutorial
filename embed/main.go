package main

import (
	"fmt"

	"m2i.fr/embed/fig"
)

func main() {
	c := fig.NewCarre(2)

	fmt.Println("Carré:", c)
	fmt.Println("Surface: ", c.GetSurface())

	arr := [5]fig.Carre{
		fig.NewCarre(2),
		fig.NewCarre(3),
		fig.NewCarre(4),
		fig.NewCarre(5),
		fig.NewCarre(6),
	}

	for i, v := range arr {
		fmt.Println("index:", i, "Valeur:", v)
	}

	slice := arr[:]
	slice2 := make([]fig.Carre, 2)

	fmt.Println("slice1:", slice, "\nslice2:", slice2)

	dict := make(map[string]fig.Carre)

	dict["carre_1"] = fig.NewCarre(2)
	dict["carre_2"] = fig.NewCarre(5)

	fmt.Println(dict)
}
