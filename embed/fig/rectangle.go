package fig

/*
 *
 */
type Rectangle struct {
	Longueur, Largeur uint
	Surface           float64
}

func (r Rectangle) GetSurface() float64 {
	return r.Surface
}

func NewRectangle(longueur, largeur uint) Rectangle {
	return Rectangle{
		Largeur: longueur, Longueur: largeur,
	}
}
