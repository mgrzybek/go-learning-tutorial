package cercle

import "math"

type Cercle struct {
	Rayon float64
}

func (c Cercle) GetSurface() float64 {
	return math.Pow(c.Rayon, 2) * math.Pi
}
