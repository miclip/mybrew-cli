package ingredients

import (
	"fmt"
	"log"

	"github.com/miclip/mybrew/hoputils"
	"github.com/miclip/mybrew/ui"
	"github.com/miclip/mybrew/utils"
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

// CreateHopInteractively adds a hop via the cli
func CreateHopInteractively(ui ui.UI) (*Hop, error) {
	h := &Hop{}
	name, err := ui.AskForText("Hop Name:")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop name with: %v", err)
	}
	h.Name = name
	amount, err := ui.AskForFloat("Amount (pounds):")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop amount with: %v", err)
	}
	h.Amount = amount
	alpha, err := ui.AskForFloat("Alpha:")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop alpha with: %v", err)
	}
	h.Alpha = alpha
	additionTime, err := ui.AskForInt("Addition Time:")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop addition time with: %v", err)
	}
	h.AdditionTime = additionTime
	form, err := ui.AskForText("Hop Form:")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop form with: %v", err)
	}
	h.Form = form
	method, err := ui.AskForText("Hop Method:")
	if err != nil {
		return nil, fmt.Errorf("Invalid hop method with: %v", err)
	}
	h.Method = method
	return h, nil
}
