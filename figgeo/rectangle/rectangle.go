package rectangle

import "errors"

// Rectangle: un polygone
type Rectangle struct {
	longueur uint
	largeur  uint
	aire     uint
}

/*
 *
 */
func (r *Rectangle) GetAire() uint {
	return r.aire
}

func updateValue(src, dest *uint) error {
	err := validateValue(src)
	if err == nil {
		*dest = *src
		return nil
	}

	return err
}

func validateValue(v *uint) error {
	if v == nil || *v == 0 {
		return errors.New("La valeur est nulle")
	}

	return nil
}

/*
 * Recalcul de l'aire si on change la taille du rectangle
 */
func (r *Rectangle) ComputeAire() {
	r.aire = r.largeur * r.longueur
}

/*
 * Changemnt des dimensions du rectangle
 * Les valeurs peuvent être nil si on ne veut changer qu'une dimension
 */
func (r *Rectangle) Resize(longueur, largeur *uint) error {
	err := updateValue(longueur, &r.longueur)
	if err != nil {
		return err
	}

	err = updateValue(largeur, &r.largeur)
	if err != nil {
		return err
	}

	r.ComputeAire()
	return nil
}

/*
 *
 */
func (r *Rectangle) GetSurface() float64 {
	return float64(r.GetAire())
}

/*
 * Factory d'un rectangle
 * Elle ne fait pas partie des méthodes de la classe mais renvoie un objet
 * instancié
 */
func New(longueur, largeur uint) (*Rectangle, error) {
	if longueur == 0 || largeur == 0 {
		return nil,
			errors.New("Un rectangle ne peut être de dimension nulle")
	}
	return &Rectangle{
		longueur: longueur,
		largeur:  largeur,
		aire:     longueur * largeur}, nil
}
