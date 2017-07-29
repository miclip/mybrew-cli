package mybrewgo

// Fermentable ...
type Fermentable struct {
	Name      string
	Amount    float64
	Yield     float64
	Potential float64
	Lovibond  float64
	Type      string
}

// Points calculates potenital yield in points
func (f *Fermentable) Points() float64 {
	return f.Yield * yieldToPoints
}

// PointsByAmount calculates potenital yield in points by amount
func (f *Fermentable) PointsByAmount() float64 {
	return f.Amount * f.Points()
}

// ColorMCU calculates color in SRM
func (f *Fermentable) ColorMCU() float64 {
	return f.Amount * f.Lovibond
}
