package mybrewgo

import "log"

// Hop ...
type Hop struct {
	Name         string
	Alpha        float64
	Amount       float64
	Form         string
	Method       string
	AdditionTime int
}

// InternationalBitteringUnits ...
func (h *Hop) InternationalBitteringUnits(hopUtils *HopUtilizations, batch float64, gravity float64) float64 {
	if h.Method != "Boil" {
		return 0.0
	}
	hopUtil, err := hopUtils.FindByAdditionTime(h.AdditionTime, gravity, h.Form)
	if err != nil {
		log.Printf("IBU's for %s could not be calculated due to error %v", h.Name, err)
		return 0.0
	}
	return (((float64(hopUtil.Percentage) / 100) * (h.Alpha / 100) * h.Amount) * 7490) / batch
}
