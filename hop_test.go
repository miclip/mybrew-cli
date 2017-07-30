package mybrewgo_test

import (
	. "github.com/miclip/mybrewgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hop", func() {
	It("Calculates International Bittering Units ", func() {
		h := Hop{
			Name:         "Galaxy",
			Amount:       1.25,
			Alpha:        13,
			Form:         "Pellet",
			Method:       "Boil",
			AdditionTime: 60,
		}
		hopUtils := NewHopUtilizations()
		ibu := h.InternationalBitteringUnits(hopUtils, 11, 1.07)
		Ω(Round(ibu, .5, 1)).Should(Equal(28.8))
	})
})
