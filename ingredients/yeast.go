package ingredients

import (
	"fmt"

	"github.com/miclip/mybrewgo/ui"
	"github.com/miclip/mybrewgo/utils"
)

// Yeast ...
type Yeast struct {
	Name        string
	Attenuation float64
}

// Print writes details of the yeasts to stdout
func (y *Yeast) Print(ui ui.UI) {
	ui.PrintLinef("%s Attenuation: %v", y.Name, utils.Round(y.Attenuation, .5, 2))
}

// CreateYeastInteractively adds a yeast via the cli
func CreateYeastInteractively(ui ui.UI) (*Yeast, error) {
	y := &Yeast{}
	name, err := ui.AskForText("Yeast Name:")
	if err != nil {
		return nil, fmt.Errorf("Invalid yeast name with: %v", err)
	}
	y.Name = name
	attenuation, err := ui.AskForFloat("Attenuation:")
	if err != nil {
		return nil, fmt.Errorf("Invalid yeast attenuation with: %v", err)
	}
	y.Attenuation = attenuation
	return y, nil
}
