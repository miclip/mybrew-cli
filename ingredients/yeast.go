package ingredients

import (
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
