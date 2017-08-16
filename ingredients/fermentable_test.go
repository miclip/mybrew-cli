package ingredients_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	. "github.com/miclip/mybrewgo/ingredients"
	"github.com/miclip/mybrewgo/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
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
	Context("UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr)
		})
		It("Prints a fermentable", func() {
			f := Fermentable{
				Name:      "2 Row",
				Amount:    23.35,
				Potential: 1.036,
				Yield:     77.9,
				Lovibond:  2,
				Type:      "Grain",
			}
			f.Print(ui)
			Ω(bOut).To(gbytes.Say("2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain"))
		})
	})
})
