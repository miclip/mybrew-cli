package ingredients

import (
	"log"

	"github.com/miclip/mybrewgo/hoputils"
	"github.com/miclip/mybrewgo/ui"
	"github.com/miclip/mybrewgo/utils"
)

// Hop represents a single hop addition to a batch
type Hop struct {
	Name         string
	Alpha        float64
	Amount       float64
	Form         string
	Method       string
	AdditionTime int
}

// InternationalBitteringUnits calculates IBU for the hop using a variation on Tinseth’s formula that
// incorporates a gravity/time adjustment instead of the bigness factor as documented by Randy Mosher
// in the "Brewer's Companion". No IBU formula is perfect so expect variations.
func (h *Hop) InternationalBitteringUnits(hopUtils *hoputils.HopUtilizations, batch float64, gravity float64) float64 {
	if h.Method != "Boil" {
		return 0.0
	}
	hopUtil, err := hopUtils.FindHopUtilization(h.AdditionTime, gravity, h.Form)
	if err != nil {
		log.Printf("IBU's for the %s addition could not be calculated due to error: %v", h.Name, err)
		return 0.0
	}
	return (((float64(hopUtil.Percentage) / 100) * (h.Alpha / 100) * h.Amount) * hopUtilOunceFactor) / batch
}

// Print writes details of the hop to stdout
func (h *Hop) Print(ui ui.UI) {
	ui.PrintLinef("%s Amount: %v Time: %v Alpha: %v Form: %s Method: %s", h.Name, utils.Round(h.Amount, .5, 2),
		h.AdditionTime, utils.Round(h.Alpha, .5, 1),
		h.Form, h.Method)
}
