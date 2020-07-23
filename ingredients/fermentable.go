package ingredients

import (
	"fmt"

	"github.com/miclip/mybrew/ui"
	"github.com/miclip/mybrew/utils"
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
func (f *Fermentable) Print(ui ui.UI) {
	ui.PrintLinef("%s Amount: %v Yield: %v Potential: %v Lovibond: %v Type: %s", f.Name, utils.Round(f.Amount, .5, 1),
		utils.Round(f.Yield, .5, 1), utils.Round(f.Potential, .5, 3), utils.Round(f.Lovibond, .5, 1), f.Type)
}

// CreateFermentableInteractively adds a fermentable via the cli
func CreateFermentableInteractively(ui ui.UI) (*Fermentable, error) {
	f := &Fermentable{}
	name, err := ui.AskForText("Fermentable Name:")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable name with: %v", err)
	}
	f.Name = name
	amount, err := ui.AskForFloat("Amount (pounds):")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable amount with: %v", err)
	}
	f.Amount = amount
	potential, err := ui.AskForFloat("Potential:")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable potential with: %v", err)
	}
	f.Potential = potential
	yield, err := ui.AskForFloat("Yield:")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable yield with: %v", err)
	}
	f.Yield = yield
	lovibond, err := ui.AskForFloat("Lovibond:")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable lovibond with: %v", err)
	}
	f.Lovibond = lovibond
	fermentableType, err := ui.AskForText("Fermentable Type:")
	if err != nil {
		return nil, fmt.Errorf("Invalid fermentable type with: %v", err)
	}
	f.Type = fermentableType
	return f, nil
}
