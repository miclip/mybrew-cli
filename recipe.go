package mybrewgo

const (
	yieldToPoints           = 0.46
	gravityBase             = 0.00001
	preBoilGallonsFactorEst = 1.2
)

// Recipe ...
type Recipe struct {
	Name               string
	Batch              float64
	BatchUnitOfMeasure string
	Style              string
	Efficiency         float64
	Method             string
	BoilTime           float64
	Hops               []Hop
	Fermentables       []Fermentable
	Yeasts             []Yeast
}

// EstimatedPreBoilVolume estimates the preboil volume
func (r *Recipe) EstimatedPreBoilVolume() float64 {
	return r.Batch * preBoilGallonsFactorEst
}

// OriginalGravity calculates the original gravity
func (r *Recipe) OriginalGravity() float64 {
	og := 0.0
	if r.Efficiency == 0 || r.Batch == 0 {
		return og
	}
	for i := range r.Fermentables {
		og = og +
			r.Fermentables[i].PointsByAmount()*(r.Efficiency*gravityBase)
	}
	return og/r.Batch + 1
}

// BoilSpecificGravity calculates the specific gravity post boil
func (r *Recipe) BoilSpecificGravity() float64 {
	return r.EstimatedPreBoilVolume()/r.Batch*(r.OriginalGravity()-1) + 1
}
