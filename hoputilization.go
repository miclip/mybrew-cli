package mybrewgo

import "fmt"

var (
	defaultPelletHopUtilizations = []HopUtilization{
		HopUtilization{5, 1.03, 5}, HopUtilization{5, 1.04, 5}, HopUtilization{5, 1.05, 4}, HopUtilization{5, 1.06, 4}, HopUtilization{5, 1.07, 3}, HopUtilization{5, 1.08, 3}, HopUtilization{5, 1.09, 3},
		HopUtilization{15, 1.03, 12}, HopUtilization{15, 1.04, 12}, HopUtilization{15, 1.05, 11}, HopUtilization{15, 1.06, 11}, HopUtilization{15, 1.07, 11}, HopUtilization{15, 1.08, 10}, HopUtilization{15, 1.09, 9},
		HopUtilization{30, 1.03, 17}, HopUtilization{30, 1.04, 17}, HopUtilization{30, 1.05, 16}, HopUtilization{30, 1.06, 16}, HopUtilization{30, 1.07, 15}, HopUtilization{30, 1.08, 15}, HopUtilization{30, 1.09, 13},
		HopUtilization{45, 1.03, 21}, HopUtilization{45, 1.04, 21}, HopUtilization{45, 1.05, 20}, HopUtilization{45, 1.06, 19}, HopUtilization{45, 1.07, 18}, HopUtilization{45, 1.08, 17}, HopUtilization{45, 1.09, 16},
		HopUtilization{60, 1.03, 24}, HopUtilization{60, 1.04, 23}, HopUtilization{60, 1.05, 23}, HopUtilization{60, 1.06, 22}, HopUtilization{60, 1.07, 21}, HopUtilization{60, 1.08, 20}, HopUtilization{60, 1.09, 18},
		HopUtilization{90, 1.03, 28}, HopUtilization{90, 1.04, 27}, HopUtilization{90, 1.05, 26}, HopUtilization{90, 1.06, 26}, HopUtilization{90, 1.07, 25}, HopUtilization{90, 1.08, 23}, HopUtilization{90, 1.09, 21},
	}
	defaultWholeHopUtilizations = []HopUtilization{
		HopUtilization{5, 1.03, 6}, HopUtilization{5, 1.04, 6}, HopUtilization{5, 1.05, 5}, HopUtilization{5, 1.06, 5}, HopUtilization{5, 1.07, 4}, HopUtilization{5, 1.08, 4}, HopUtilization{5, 1.09, 3},
		HopUtilization{15, 1.03, 15}, HopUtilization{15, 1.04, 15}, HopUtilization{15, 1.05, 14}, HopUtilization{15, 1.06, 14}, HopUtilization{15, 1.07, 13}, HopUtilization{15, 1.08, 13}, HopUtilization{15, 1.09, 11},
		HopUtilization{30, 1.03, 22}, HopUtilization{30, 1.04, 21}, HopUtilization{30, 1.05, 21}, HopUtilization{30, 1.06, 20}, HopUtilization{30, 1.07, 19}, HopUtilization{30, 1.08, 18}, HopUtilization{30, 1.09, 16},
		HopUtilization{45, 1.03, 26}, HopUtilization{45, 1.04, 26}, HopUtilization{45, 1.05, 25}, HopUtilization{45, 1.06, 24}, HopUtilization{45, 1.07, 23}, HopUtilization{45, 1.08, 22}, HopUtilization{45, 1.09, 21},
		HopUtilization{60, 1.03, 29}, HopUtilization{60, 1.04, 28}, HopUtilization{60, 1.05, 28}, HopUtilization{60, 1.06, 27}, HopUtilization{60, 1.07, 26}, HopUtilization{60, 1.08, 25}, HopUtilization{60, 1.09, 23},
		HopUtilization{90, 1.03, 35}, HopUtilization{90, 1.04, 34}, HopUtilization{90, 1.05, 33}, HopUtilization{90, 1.06, 32}, HopUtilization{90, 1.07, 31}, HopUtilization{90, 1.08, 29}, HopUtilization{90, 1.09, 27},
	}
)

// HopUtilizations ...
type HopUtilizations struct {
	Matrix map[string][]HopUtilization
}

// HopUtilization ...
type HopUtilization struct {
	Minutes    int
	Gravity    float64
	Percentage int
}

// NewHopUtilizations creates a new HopUtilization matrix
func NewHopUtilizations() HopUtilizations {
	h := make(map[string][]HopUtilization)
	h["Pellet"] = defaultPelletHopUtilizations
	h["Whole"] = defaultWholeHopUtilizations
	return HopUtilizations{
		Matrix: h,
	}
}

// FindByAdditionTime ...
func (h *HopUtilizations) FindByAdditionTime(additionTime int, gravity float64, hopType string) (lower *HopUtilization, upper *HopUtilization, err error) {
	hopUtils := h.Matrix[hopType]
	var l, u *HopUtilization
	for i := range hopUtils {
		if hopUtils[i].Minutes == h.findHopLowerMinutes(additionTime) && Round(hopUtils[i].Gravity, .5, 2) == Round(gravity, .5, 2) {
			l = &hopUtils[i]
		}
		if hopUtils[i].Minutes == h.findHopUpperMinutes(additionTime) && Round(hopUtils[i].Gravity, .5, 2) == Round(gravity, .5, 2) {
			u = &hopUtils[i]
		}
	}
	if l != nil && u != nil {
		return l, u, nil
	}

	return nil, nil, fmt.Errorf("No HopUtilization was found for AdditionTime: %v Gravity: %v", additionTime, gravity)
}

func (h *HopUtilizations) findHopUpperMinutes(hopTime int) int {
	mins := []int{5, 15, 30, 45, 60, 90}
	for i := range mins {
		if mins[i] > hopTime {
			return mins[i]
		}
	}
	return 90
}

func (h *HopUtilizations) findHopLowerMinutes(hopTime int) int {
	mins := []int{90, 60, 45, 30, 15, 5}
	for i := range mins {
		if mins[i] <= hopTime {
			return mins[i]
		}
	}
	return 5
}
