package ingredients_test

import (
	. "github.com/miclip/mybrewgo/ingredients"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fermentable", func() {
	Context("Fermentable Calculations", func() {
		var f Fermentable
		BeforeEach(func() {
			f = Fermentable{
				Name:      "2 Row",
				Amount:    23.35,
				Potential: 1.036,
				Yield:     77.9,
				Lovibond:  2,
			}
		})
		It("Calculates the Points()", func() {
			Ω(f.Points()).Should(Equal(35.834))
		})
		It("Calculates the PointsByAmount()", func() {
			Ω(f.PointsByAmount()).Should(Equal(836.7239000000001))
		})
		It("Calculates the ColorMCU()", func() {
			Ω(f.ColorMCU()).Should(Equal(46.7))
		})
	})

})
