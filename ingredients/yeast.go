package ingredients

import (
	"github.com/fatih/color"
	"github.com/miclip/mybrewgo/utils"
)

// Yeast ...
type Yeast struct {
	Name        string
	Attenuation float64
}

// Print writes details of the yeasts to stdout
func (y *Yeast) Print() {
	color.Magenta("%s Attenution: %v", y.Name, utils.Round(y.Attenuation, .5, 2))
}
