package ingredients

import (
	"github.com/fatih/color"
	"github.com/miclip/mybrewgo/utils"
)

// Fermentable ...
type Fermentable struct {
	Name      string
	Amount    float64
	Yield     float64
	Potential float64
	Lovibond  float64
	Type      string
}

const (
	yieldToPoints           = 0.46
	gravityBase             = 0.00001
	preBoilGallonsFactorEst = 1.2
	hopUtilOunceFactor      = 7490
)

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

// Print writes details of the fermentable to stdout
func (f *Fermentable) Print() {
	color.Yellow("%s Amount: %v Yield: %v Potential: %v Lovibond: %v Type: %s", f.Name, utils.Round(f.Amount, .5, 1),
		utils.Round(f.Yield, .5, 1), utils.Round(f.Potential, .5, 3), utils.Round(f.Lovibond, .5, 1), f.Type)
}
