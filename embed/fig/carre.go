package fig

import (
	"fmt"
	"reflect"
)

type Carre struct {
	Rectangle
	Cote uint
}

func (c Carre) GetSurface() float64 {
	return c.Surface
}

func NewCarre(cote uint) Carre {
	surface := float64(cote * cote)
	return Carre{
		Cote: cote,
		Rectangle: Rectangle{
			Largeur: cote, Longueur: cote, Surface: surface,
		},
	}
}

func (c Carre) String() string {
	return fmt.Sprintf("Type: %v, Valeur: %v", reflect.TypeOf(c), c.Cote)
}
